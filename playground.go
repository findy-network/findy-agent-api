package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/findy-network/findy-agent-api/server"
	"github.com/findy-network/findy-agent-api/tools"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const defaultPort = "8085"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// TEST subscription
	ticker := time.NewTicker(time.Second * 20)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				tools.AddEvent()
			}
		}
	}()
	// TEST SUBSCRIPTION end

	srv := server.Server(&tools.Resolver{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// TODO:
				return true
			},
		},
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.AllowAll().Handler(srv)) // TODO: adjust CORS policy

	// TODO: adjust CORS policy
	//http.Handle("/", cors.AllowAll().Handler(r))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
