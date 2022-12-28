package main

import (
	"github.com/ducdang91/go-gqlgen-dataloader/config"
	"github.com/ducdang91/go-gqlgen-dataloader/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ducdang91/go-gqlgen-dataloader/migration"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"


func init() {
	// Load Environment from .env
	godotenv.Load()

	// Initialize db Value
	config.ConnectGorm()
}

func main() {
	// Defer close database connection (optional)
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Migration
	migration.MigrateTable()
	
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