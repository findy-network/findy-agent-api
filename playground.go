package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/findy-network/findy-agent-api/server"
	"github.com/findy-network/findy-agent-api/tools"
)

const defaultPort = "8085"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", server.Handler(&resolver.Resolver{}))
	http.Handle("/query", server.Handler(&tools.Resolver{}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
