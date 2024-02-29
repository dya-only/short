package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Short struct {
	gorm.Model
	Key string `gorm:"unique"`
	Url string
}

type Request struct {
	Url string `json:"url"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getRandKey(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func main() {

	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Short{})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		// Return Short Url
		if r.Method == http.MethodGet {
			// key := strings.TrimPrefix(r.URL.Path, "/api")
			key := r.URL.Query().Get("key")

			var short Short
			db.First(&short, "key = ?", key)

			resp := make(map[string]string)
			resp["url"] = short.Url
			JSONresp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(JSONresp)
			return
		}

		// Create New Short Url
		if r.Method == http.MethodPost {
			var body Request
			err := json.NewDecoder(r.Body).Decode(&body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			randKey := getRandKey(10)
			db.Create(&Short{Key: randKey, Url: body.Url})

			resp := make(map[string]string)
			resp["key"] = randKey
			JSONresp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(JSONresp)
			return
		}

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	http.ListenAndServe(":8080", nil)
}
