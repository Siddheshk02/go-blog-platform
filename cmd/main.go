package main

import (
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
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres", "host="+dbHost+" port="+dbPort+" user="+dbUser+" dbname="+dbName+" password="+dbPassword+" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Post{})
	authRepo := repository.NewAuthRepository(db)
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

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
