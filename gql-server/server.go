package main

import (
	"fmt"
	"gql-server/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// Database
	"context"
	"database/sql"
	"gql-server/database"
	"reflect"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	if err := run(); err != nil {
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

var dataSource string = "database/todos.db"

func run() error {
	ctx := context.Background()

	conn, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return err
	}
	fmt.Printf("Connected to %s\n", dataSource)

	db := database.New(conn)

	// 2. list all todos
	todos, err := db.GetTodos(ctx)
	if err != nil {
		return err
	}
	log.Println(todos)

	// 3. create a todo
	insertedTodo, err := db.CreateTodo(ctx, "Mama Mia!")
	
	if err != nil {
		return err
	}
	log.Println(insertedTodo)

	// 4. get the todo we just inserted
	fetchedTodo, err := db.GetTodo(ctx, insertedTodo.ID)
	if err != nil {
		return err
	}

	// 5. prints true
	log.Println(reflect.DeepEqual(insertedTodo, fetchedTodo))
	return nil
}