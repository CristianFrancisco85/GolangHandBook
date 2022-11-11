package handlers

import (
	"encoding/json"
	"net/http"
	"rest_ws/models"
	"rest_ws/repository"
	"rest_ws/server"
	"rest_ws/utils"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

type UpdateInsertPostRequest struct {
	Content string `json:"content"`
}

type InsertPostResponse struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type UpdatePostResponse struct {
	Message string `json:"message"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetTokenFromHeader(r, s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := utils.GetUserIdFromToken(r, token)
		if err != nil || user == nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		var request = UpdateInsertPostRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post := models.Post{
			Id:      id.String(),
			Content: request.Content,
			UserID:  user.Id,
		}

		err = repository.InsertPost(r.Context(), &post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postMessage = models.WebSocketMessage{
			Type:    "post",
			Payload: post,
		}
		s.Hub().Broadcast(postMessage, nil)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(InsertPostResponse{
			Id:      post.Id,
			Content: post.Content,
		})

	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "Invalid Id", http.StatusBadRequest)
			return
		}

		post, err := repository.GetPostById(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if post == nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := utils.GetTokenFromHeader(r, s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := utils.GetUserIdFromToken(r, token)
		if err != nil || user == nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "Invalid Id", http.StatusBadRequest)
			return
		}

		post, err := repository.GetPostById(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if post == nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		if post.UserID != user.Id {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var request = UpdateInsertPostRequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		post.Content = request.Content

		err = repository.UpdatePost(r.Context(), post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UpdatePostResponse{
			Message: "Post updated",
		})
	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := utils.GetTokenFromHeader(r, s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := utils.GetUserIdFromToken(r, token)
		if err != nil || user == nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "Invalid Id", http.StatusBadRequest)
			return
		}

		post, err := repository.GetPostById(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if post == nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		if post.UserID != user.Id {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err = repository.DeletePost(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UpdatePostResponse{
			Message: "Post deleted",
		})
	}
}

func ListPostsHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := utils.GetTokenFromHeader(r, s.Config().JWTSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := utils.GetUserIdFromToken(r, token)
		if err != nil || user == nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		pages, err := strconv.ParseUint(mux.Vars(r)["pages"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid Pages", http.StatusBadRequest)
			return
		}
		posts, err := repository.ListPosts(r.Context(), pages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
	}
}
