package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/findy-network/findy-agent-api/server"
	"github.com/findy-network/findy-agent-api/tools/resolver"
	"github.com/rs/cors"
)

const defaultPort = "8085"

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		glog.V(3).Infof("received request: %s %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}

func initLogging() {
	flag.Set("logtostderr", "true")
	flag.Set("stderrthreshold", "WARNING")
	flag.Set("v", "3")
	flag.Parse()
}

func main() {
	initLogging()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	/*// TEST subscription
	ticker := time.NewTicker(time.Second * 30)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				resolver.AddEvent()
			}
		}
	}()
	// TEST SUBSCRIPTION end*/

	srv := server.Server(&resolver.Resolver{})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.AllowAll().Handler(logRequest(srv)))

	glog.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
