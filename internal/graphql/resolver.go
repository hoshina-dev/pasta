package graphql

import "github.com/hoshina-dev/pasta/internal/service"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	PastaService *service.PastaService
}
