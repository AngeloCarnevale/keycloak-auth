package routes

import (
	"demo-backend/api/handlers"
	"demo-backend/infrastructure/identity"
	usemanagementusercase "demo-backend/use_cases/use_management_usecase"

	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to My Demo Rest API"))
	})

	grp := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usemanagementusercase.NewRegisterUseCase(identityManager)
	
	grp.Post("/user", handlers.RegisterHandler(registerUseCase))
}