package user

import (
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func ApplyRoutes(r fiber.Router) {
	apis := r.Group("user", middleware.CheckStore())
	{
		apis.Post("", CreateUser)
		apis.Get("/:userName", GetUser)
		//TODO  μλ api μΆκ°
		//apis.Put("", UpdateUser)
		//apis.Delete("", DeleteUser)
	}
}
