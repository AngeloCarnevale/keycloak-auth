package middlewares

import (
	"context"
	"demo-backend/shared/enums"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitFiberMiddleware(app *fiber.App,
	initPublicRoutes func(app *fiber.App),
	initProtectedRoutes func(app *fiber.App)) {
	
	app.Use(requestid.New())
	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		var requestId = c.Locals("requestid")

		var ctx = context.WithValue(context.Background(),enums.ContextKeyRequestId, requestId)
		c.SetUserContext(ctx)

		return c.Next()
	})

	initPublicRoutes(app)
	log.Println("fiber middlewares initialized")
}