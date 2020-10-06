package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/findy-network/findy-agent-api/graph/generated"
)

func schema(resolver generated.ResolverRoot) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{Resolvers: resolver})
}

// Handler provides graphQL query handler
func Server(resolver generated.ResolverRoot) *handler.Server {
	srv := handler.NewDefaultServer(schema(resolver))
	srv.AddTransport(transport.POST{})

	return srv
}
