package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("[SERVER] Database initialization failed: %v\n", err)
	}
	defer db.Close()

	http.HandleFunc("/", handlePost(db))

	log.Println("[SERVER] Starting HTTP server on :1234")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalf("[SERVER] Failed to start server: %v\n", err)
	}
}

func initDB() (*sql.DB, error) {
	// https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-parameters
	db, err := sql.Open("mysql", "ids_project_2024:ids_project_2024@tcp(db:3306)/Articles?multiStatements=true")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func handlePost(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == "POST" {
			var post Post
			if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Safe: _, err := db.Exec("INSERT INTO Posts (title, content) VALUES (?, ?)", post.Title, post.Content)
			// see: https://go.dev/doc/database/sql-injection
			payload := fmt.Sprintf("INSERT INTO Posts (title, content) VALUES ('%s', '%s')", post.Title, post.Content)
			log.Println(payload)
			_, err := db.Exec(payload)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}

		if r.Method == "GET" {
			rows, err := db.Query("SELECT title, content FROM Posts")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var posts []Post
			for rows.Next() {
				var post Post
				if err := rows.Scan(&post.Title, &post.Content); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				posts = append(posts, post)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
