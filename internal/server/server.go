package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hoshina-dev/pasta/internal/graphql"
	webhookHandler "github.com/hoshina-dev/pasta/internal/handler"
)

func New(resolver *graphql.Resolver, webhook *webhookHandler.WebhookHandler, corsOrigins string) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: corsOrigins,
	}))

	srv := handler.New(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: resolver,
	}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	app.All("/graphql", adaptor.HTTPHandler(srv))
	app.Get("/", adaptor.HTTPHandler(
		playground.Handler("Pasta GraphQL", "/graphql"),
	))
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Post("/webhook/optimization", webhook.HandleOptimizationCallback)

	return app
}
