package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/haedarrfd/simple-rss-aggregator/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// Hold a connection to database (in this case postgres)
type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Setting up the environment
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env")
	}

	// Open a connection to postgresql
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	// apiConfig to pass database handler, so they have access to database
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// Set up an HTTP server and listen on the given port
	router := chi.NewRouter()

	// Set up CORS (Cross-Origin Resource Sharing) middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create sub-path route, then the route is handled by handler function
	v1Router := chi.NewRouter()
	v1Router.Get("/harmony", handlerRead)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)

	// Mount this router on the main router path
	router.Mount("/v1", v1Router)

	// Create a new HTTP server configuration
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v\n", portString)

	// Start the server, if any error appear immediately stop the program
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
