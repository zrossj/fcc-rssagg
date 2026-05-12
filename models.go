package main

import (
	"time"

	"gibhub.com/jzross/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
}

func returnPOST(u database.User) User {

	return User{
		ID:         u.ID,
		CreatedAt:  u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		Name:       u.Name,
	}
}
