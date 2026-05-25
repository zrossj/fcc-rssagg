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
	ApiKey     string    `json:"api_key"`
}

func returnPOST(u database.User) User {

	return User{
		ID:         u.ID,
		CreatedAt:  u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		Name:       u.Name,
		ApiKey:     u.ApiKey,
	}
}

type Feed struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"crated_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Url        string    `json:"url"`
	UserID     uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(p database.Feed) Feed {
	return Feed{
		ID:         p.ID,
		CreatedAt:  p.CreatedAt,
		ModifiedAt: p.ModifiedAt,
		Name:       p.Name,
		Url:        p.Url,
		UserID:     p.UserID}
}

func databaseFeedsToFeeds(f []database.Feed) []Feed {

	var vals []Feed
	for _, feed := range f {
		vals = append(vals, databaseFeedToFeed(feed))
	}

	return vals
}

type SubFeed struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	UserID     uuid.UUID `json:"user_id"`
	FeedID     uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(fs database.Feedsub) SubFeed {
	return SubFeed{
		ID:         fs.ID,
		CreatedAt:  fs.CreatedAt,
		ModifiedAt: fs.ModifiedAt,
		UserID:     fs.UserID,
		FeedID:     fs.FeedID,
	}
}

func databaseUserFeedsToUserFeeds(ufeeds []database.Feedsub) []SubFeed {

	vals := []SubFeed{}
	for _, feedsub := range ufeeds {
		vals = append(vals, databaseFeedFollowToFeedFollow(feedsub))
	}
	return vals
}

type GetLastPosts struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPosts(posts []database.Post) []GetLastPosts {

	var formatedPosts []GetLastPosts
	var descrip *string

	for _, post := range posts {

		if post.Description.Valid {
			descrip = &post.Description.String
		}

		formatedPosts = append(
			formatedPosts,
			GetLastPosts{
				ID:          post.ID,
				CreatedAt:   post.CreatedAt,
				UpdatedAt:   post.UpdatedAt,
				Title:       post.Title,
				Description: descrip,
				PublishedAt: post.PublishedAt,
				Url:         post.Url,
				FeedID:      post.FeedID,
			})
	}
	return formatedPosts
}
