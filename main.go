package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	fmt.Println("Hello Word")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	db := database.New(conn)

	apiCfg := apiConfig{
		DB: db,
	}

	fmt.Println("Port: ", portString)

	go ScraperTriggerRoutine(db, time.Duration(30*time.Minute), 10)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/Err", handlerErr)
	v1Router.Post("/createUser", apiCfg.handlerCreateUser)
	v1Router.Get("/createUser", apiCfg.middleware(apiCfg.handlerGetUser))
	v1Router.Post("/createFeed", apiCfg.middleware(apiCfg.handleCreateFeed))
	v1Router.Get("/createFeed", apiCfg.getFeeds)

	v1Router.Post("/subscribeFeed", apiCfg.middleware(apiCfg.handleSubsToFeed))
	v1Router.Get("/subscribeFeed", apiCfg.middleware(apiCfg.handlerGetUserFeeds))
	v1Router.Delete("/subscribeFeed/{FeedID}", apiCfg.middleware(apiCfg.handleDeleteSub))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port: %v", portString)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
