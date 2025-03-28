package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PostCount int    `json:"postCount"`
}

// Post struct
type Post struct {
	ID           int    `json:"id"`
	User         string `json:"user"`
	Content      string `json:"content"`
	CommentCount int    `json:"commentCount"`
	Timestamp    int64  `json:"timestamp"`
}

// In-memory data
var users = []User{
	{1, "Ayush", 10}, {2, "Mohan", 7}, {3, "Sujoy", 15},
	{4, "Shrey", 8}, {5, "Raavan", 12}, {6, "Mahi", 5},
	{7, "Ravi", 20}, {8, "Suresh", 3}, {9, "Kumar", 6},
}

var posts = []Post{
	{1, "Ayush", "Hello World!", 20, 1711500},
	{2, "Mohan", "Go is awesome!", 15, 1811510000},
	{3, "Sujoy", "Next.js and Go!", 30, 1711520000},
	{4, "Shrey", "Trending post!", 30, 1211530000},
	{5, "Raavan", "Learning Fiber!", 10, 1511540000},
	{6, "Mahi", "Golang is great!", 5, 1611550000},
	{7, "Ravi", "Backend development!", 25, 1711560000},
}

// Middleware for CORS
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Function to get top 5 users using a max heap approach
func getTopUsers(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Find top 5 users using a simple selection method
	topUsers := make([]User, 5)
	for i := 0; i < 5; i++ {
		maxIdx := i
		for j := i + 1; j < len(users); j++ {
			if users[j].PostCount > users[maxIdx].PostCount {
				maxIdx = j
			}
		}
		users[i], users[maxIdx] = users[maxIdx], users[i]
		topUsers[i] = users[i]
	}

	json.NewEncoder(w).Encode(topUsers)
}

// Function to get trending (most commented) or latest posts
func getPosts(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	query := r.URL.Query().Get("type")

	if query == "popular" {
		// Find max comment count
		maxComments := 0
		for _, post := range posts {
			if post.CommentCount > maxComments {
				maxComments = post.CommentCount
			}
		}

		// Collect all posts with the max comment count
		var trendingPosts []Post
		for _, post := range posts {
			if post.CommentCount == maxComments {
				trendingPosts = append(trendingPosts, post)
			}
		}
		json.NewEncoder(w).Encode(trendingPosts)
		return
	}

	// Default: Get latest 5 posts using selection sort (efficient for small data)
	latestPosts := make([]Post, 5)
	for i := 0; i < 5; i++ {
		maxIdx := i
		for j := i + 1; j < len(posts); j++ {
			if posts[j].Timestamp > posts[maxIdx].Timestamp {
				maxIdx = j
			}
		}
		posts[i], posts[maxIdx] = posts[maxIdx], posts[i]
		latestPosts[i] = posts[i]
	}

	json.NewEncoder(w).Encode(latestPosts)
}

func liveFeed(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

}

func main() {
	http.HandleFunc("/users", getTopUsers) // mapping functions to routes
	http.HandleFunc("/posts", getPosts)
	http.HandleFunc("/livefeed", liveFeed)

	log.Println("Go Backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
