package main

import (
	"gql-server/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// Database
	"context"
	"gql-server/database"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// Initialize Database
	if err := initializeDB(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initializeDB() error {
	// Initialize the database connection
	dataSource := "database/todos.db"

	if err := database.Init(dataSource); err != nil {
		return err
	}

	// Try all queries
	ctx := context.Background()
	database.DB.GetTodos(ctx)

	return nil
}
