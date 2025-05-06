package response

import "github.com/gofiber/fiber/v2"

func ErrorResponse(ctx *fiber.Ctx, status int, message string, err error) error {
	return ctx.Status(status).JSON(&fiber.Map{
		"error": message + ": " + err.Error(),
		"data":  nil,
	})
}

func SuccessResponse(ctx *fiber.Ctx, status int, message string, data interface{}) error {
	return ctx.Status(status).JSON(&fiber.Map{
		"success": message,
		"data":    data,
	})
}
