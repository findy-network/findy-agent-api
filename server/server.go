package server

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/findy-network/findy-agent-api/graph/generated"
)

// Handler provides graphQL query handler
func Handler(resolver generated.ResolverRoot) http.Handler {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
}
