package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mychell/go-rest-starter/driver"
	"github.com/mychell/go-rest-starter/models"
	"github.com/mychell/go-rest-starter/repository"
	"github.com/mychell/go-rest-starter/repository/post"
)

// Post ...
type Post struct {
	repository repository.PostRepository
}

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repository: post.NewPQPostRepository(db.PQ),
	}
}

// Create a new post
func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	post := models.POST{}
	json.NewDecoder(r.Body).Decode(&post)

	err := p.repository.Create(&post)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
