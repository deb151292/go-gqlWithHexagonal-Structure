package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	database "travel.go/adoptors/output/detabase"
	graph "travel.go/gql/generated"
	resolve "travel.go/gql/resolver"
)
const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading env", err)
	}

	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = defaultPort
	}

	database.DatabaseConn()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolve.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
