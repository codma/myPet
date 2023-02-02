package middleware

import (
	"main/database/repositories"
	"main/model"

	"github.com/gofiber/fiber/v2"
)

func CheckStore() fiber.Handler {
	return func(c *fiber.Ctx) error {

		storeId := c.Get("X-Sixshop-Metadata-Storeid")
		token := c.Get("Authorization-App")

		err := repositories.ValidationAppKey(storeId, token)
		if err != nil {
			c.Response().SetStatusCode(400)
			c.JSON(model.FailResponse{
				Message: "스토어를 검증하는데 실패했습니다.",
				Error:   err,
			})
		}

		c.Locals("storeId", storeId)
		return c.Next()
	}
}
