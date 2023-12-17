package handlers

import (
	"context"
	productuc "demo-backend/use_cases/product_uc"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, request productuc.CreateProductRequest) (*productuc.CreateProductResponse, error)
}

func CreateProductHandler(useCase CreateProductUseCase) fiber.Handler{
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()

		var request = productuc.CreateProductRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.CreateProduct(ctx, request)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}