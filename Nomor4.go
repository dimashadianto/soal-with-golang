package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Author      string `json:"author"`
	NumComments int    `json:"num_comments"`
	StoryID     interface{} `json:"story_id"`
	StoryTitle  interface{} `json:"story_title"`
	StoryURL    interface{} `json:"story_url"`
	ParentID    interface{} `json:"parent_id"`
	CreatedAt   int    `json:"created_at"`
}

type Articles struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Article `json:"data"`
}

func getArticles(username, page string) {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?author=%s&page=%s", username, page)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var responseObject Articles
		err := json.NewDecoder(response.Body).Decode(&responseObject)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		for _, rec := range responseObject.Data {
			if rec.Title != "" {
				fmt.Println(rec.Title)
			} else if rec.StoryTitle != nil {
				fmt.Println(rec.StoryTitle)
			}

		}
	} else {
		fmt.Println("HTTP request failed:", response.Status)
	}
}

func main() {
	getArticles("olalonde", "1")
}