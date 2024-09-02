package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Siddheshk02/go-blog-platform/config"
	"github.com/Siddheshk02/go-blog-platform/controllers"
	"github.com/Siddheshk02/go-blog-platform/models"
	repository "github.com/Siddheshk02/go-blog-platform/repositories"
	"github.com/Siddheshk02/go-blog-platform/routes"
	"github.com/Siddheshk02/go-blog-platform/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	config.LoadConfig()
	user := os.Getenv("User")
	pass := os.Getenv("Password")

	connStr := "host=localhost port=5432 user=" + user + " password=" + pass + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	databaseName := "go_blog_platform"
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Printf("Database %s created successfully", databaseName)
	}

	gormConnStr := fmt.Sprintf("%s dbname=%s", connStr, databaseName)
	gormDB, err := gorm.Open("postgres", gormConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to the new database with Gorm: %v", err)
	}
	defer gormDB.Close()

	gormDB.AutoMigrate(&models.User{}, &models.Post{})

	authRepo := repository.NewAuthRepository(gormDB)
	userRepo := repository.NewUserRepository(gormDB)
	postRepo := repository.NewPostRepository(gormDB)

	authService := services.NewAuthService(authRepo)
	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)

	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	postController := controllers.NewPostController(postService)

	router := routes.SetupRoutes(authController, userController, postController)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
