# Go Blog Platform

A simple blog platform API built with Golang featuring JWT authentication, CRUD operations for posts, and user management.

## Getting Started

### Prerequisites

- Go 1.20 or higher
- PostgreSQL
- Docker 

### Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/your-repo/go-blog-platform.git
    cd go-blog-platform
    ```

2. Run the application:

    ```bash
    docker pull siddheshk02/go-blog-platform:latest
    ```
    ```bash
    docker compose up -d
    ```
On Successful Run, test the API (Use any API testing tool)

3. Access the API at `http://localhost:8080`.

## Example
 - Register : (POST) `http://localhost:8080/register` { "username": "your_username", "email": "your_email", "password": "your_password" }
   
 - Login : (POST) `http://localhost:8080/login` { "username": "your_username", "email": "your_email", "password": "your_password" } ("When you log in, the API generates a JWT bearer token, which can be found in the response header under Authorization. Use this token by adding it to the Authorization header of your subsequent requests (e.g., Authorization: Bearer <token>) to access protected routes.")

 - Profile : (GET) `http://localhost:8080/profile` {"user_id":1,"username":"User1","email":"User1@xyz.com"}

 - Profile : (GET) `http://localhost:8080/profile/{id}` (Get a specific profile using id)

 - Profile : (PUT) `http://localhost:8080/profile/{id}` {"username":"Updated username","email":"Updated Email"}

 - Posts : (POST) `http://localhost:8080/posts` {"title":"Blog Title","content":"Simple Tutorial Blog"}

 - Posts : (GET) `http://localhost:8080/posts` (Get all posts)

 - Posts : (GET) `http://localhost:8080/posts/{id}` {"ID":1,"CreatedAt":"2024-09-02T07:59:04.04957Z","UpdatedAt":"2024-09-02T07:59:34.16749Z","DeletedAt":null,"title":"WebRTC in Go","content":"WebRTC Basic Tutorial Blog Post","user_id":1}

 - Posts : (PUT) `http://localhost:8080/posts/{id}` {"title":"Update Title","content":"Updated Content"}

 - Posts : (DELETE) `http://localhost:8080/posts/{id}` (Delete a posts using id)

## API Documentation

For detailed API endpoints and usage, refer to the `swagger.yaml` file in the `docs/` folder.

