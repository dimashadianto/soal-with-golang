package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type Article struct {
	Title       *string `json:"title"`
	URL         string  `json:"url"`
	Author      string  `json:"author"`
	NumComments int     `json:"num_comments"`
	StoryID     *int    `json:"story_id"`
	StoryTitle  *string `json:"story_title"`
	StoryURL    *string `json:"story_url"`
	ParentID    *int    `json:"parent_id"`
	CreatedAt   int     `json:"created_at"`
}

type Articles struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	Total      int       `json:"total"`
	TotalPages int       `json:"total_pages"`
	Data       []Article `json:"data"`
}

func getArticles(username, page string) {
	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles")
	if username != "" && page != "" {
		url += fmt.Sprintf("?author=%s&page=%s", username, page)
	}

	if username != "" && page == "" {
		url += fmt.Sprintf("?author=%s", username)
	}

	if username == "" && page != "" {
		url += fmt.Sprintf("?page=%s", page)
	}

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

		result := []string{}
		for _, rec := range responseObject.Data {
			if rec.Title != nil {
				if *rec.Title != "" {
					result = append(result, *rec.Title)
				}
			} else if rec.StoryTitle != nil {
				if *rec.StoryTitle != "" {
					result = append(result, *rec.StoryTitle)
				}
			}
		}
		fmt.Println("Result:", result)

	} else {
		fmt.Println("HTTP request failed:", response.Status)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ClearScreen()
	fmt.Print("Input Author :")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Input Page :")
	scanner.Scan()
	id := scanner.Text()
	getArticles(name, id)
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Untuk Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
