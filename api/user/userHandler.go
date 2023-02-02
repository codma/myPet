package user

import (
	"main/database/repositories"
	"main/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) (err error) {
	var user model.User
	err = c.BodyParser(&user)
	if err != nil {
		c.Response().SetStatusCode(400)
		c.JSON(model.FailResponse{
			Message: "데이터를 읽는데 실패했습니다.",
			Error:   err,
		})
		return
	}

	storeId := c.Locals("storeId").(string)
	memberId, err := repositories.CreateUser(user, storeId)
	if err != nil {
		c.Response().SetStatusCode(400)
		c.JSON(err.Error())
		return
	}

	response := model.CreateUserResponse{
		MemberId: memberId,
	}

	c.Response().SetStatusCode(200)
	c.JSON(response)
	return
}

func GetUser(c *fiber.Ctx) (err error) {
	userName := c.Params("userName", "")
	storeId := c.Locals("storeId").(string)

	if userName == "" {
		c.Response().SetStatusCode(400)
		c.JSON(model.FailResponse{
			Message: "userId를 확인하는데 실패했습니다.",
		})
		return
	}

	user, err := repositories.GetUser(userName, storeId)
	if err != nil {
		c.Response().SetStatusCode(400)
		c.JSON(err.Error())
		return
	}

	c.Response().SetStatusCode(200)
	c.JSON(user)
	return
}
