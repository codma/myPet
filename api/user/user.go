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
		//TODO  아래 api 추가
		//apis.Put("", UpdateUser)
		//apis.Delete("", DeleteUser)
	}
}
