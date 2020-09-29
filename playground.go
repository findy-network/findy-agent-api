package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/findy-network/findy-agent-api/server"
	"github.com/findy-network/findy-agent-api/tools"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const defaultPort = "8085"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()
	// TODO: adjust CORS policy
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	r.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	r.HandleFunc("/query", server.Handler(&tools.Resolver{}).ServeHTTP)
	http.Handle("/", c.Handler(r))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
