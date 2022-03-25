package main

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Abdubek/auth-test/graph/service"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Abdubek/auth-test/graph"
	"github.com/Abdubek/auth-test/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	c := cors.AllowAll()
	router.Use(c.Handler)
	router.Use(service.Middleware())

	gqlConfig := generated.Config{
		Resolvers: &graph.Resolver{},
	}
	gqlConfig.Directives.AuthGuard = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		authDetails := service.ForContext(ctx)
		if authDetails == nil {
			return nil, errors.New("INVALID_TOKEN")
		}
		return next(ctx)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))

	router.Handle("/frontend/task", playground.Handler("GraphQL playground", "/frontend/task/graphql"))
	router.Handle("/frontend/task/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
