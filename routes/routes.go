package routes

import (
	"github.com/Siddheshk02/go-blog-platform/controllers"
	"github.com/Siddheshk02/go-blog-platform/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes(authController *controllers.AuthController, userController *controllers.UserController, postController *controllers.PostController) *mux.Router {

	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/register", authController.RegisterUser).Methods("POST")
	router.HandleFunc("/login", authController.Login).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.JWTAuth)

	// User routes
	protected.HandleFunc("/profiles", userController.GetAllUsers).Methods("GET")
	protected.HandleFunc("/profile/{id}", userController.GetUserProfile).Methods("GET")
	protected.HandleFunc("/profile/{id}", userController.UpdateUser).Methods("PUT")

	// Post routes
	protected.HandleFunc("/posts", postController.CreatePost).Methods("POST")
	protected.HandleFunc("/posts", postController.GetAllPosts).Methods("GET")
	protected.HandleFunc("/posts/{id}", postController.GetPost).Methods("GET")
	protected.HandleFunc("/posts/{id}", postController.UpdatePost).Methods("PUT")
	protected.HandleFunc("/posts/{id}", postController.DeletePost).Methods("DELETE")

	return router
}
