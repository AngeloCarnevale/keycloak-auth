package handlers

import (
	"context"
	"demo-backend/use_cases/use_management_usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type RegisterCase interface {
	Register(context.Context, usemanagementusercase.RegisterRequest) (*usemanagementusercase.RegisterResponse, error)
}

func RegisterHandler(useCase RegisterCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		var request = usemanagementusercase.RegisterRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.Register(ctx, request)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}