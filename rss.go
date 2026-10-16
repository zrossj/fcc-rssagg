package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func getRSSFeed(url string) (RSSFeed, error) {

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	fmt.Printf("URL being scraped: %v\n", url)
	response, err := client.Get(url)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to request url: %w", err)
	}

	defer response.Body.Close()

	feedBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to read response body: %w", err)
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(feedBytes, &feed)

	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to parse RSS XML: %w", err)
	}

	return feed, nil
}
