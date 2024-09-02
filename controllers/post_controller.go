package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Siddheshk02/go-blog-platform/models"
	"github.com/Siddheshk02/go-blog-platform/services"
	"github.com/gorilla/mux"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{postService: postService}
}

func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("userID").(uint)
	post.UserID = userID

	err = c.postService.CreatePost(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Post ID", http.StatusBadRequest)
		return
	}

	post, err := c.postService.FetchPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}

func (c *PostController) GetAllPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := c.postService.FetchAllPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func (c *PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	var postIn models.PostInput
	if err := json.NewDecoder(r.Body).Decode(&postIn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var post models.Post
	post.Title = postIn.Title
	post.Content = postIn.Content
	err = c.postService.UpdatePost(id, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	err = c.postService.DeletePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
