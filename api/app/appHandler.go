package app

import (
	"main/common"
	"main/model"

	"github.com/gofiber/fiber/v2"
)

func CreateOrderForTest(c *fiber.Ctx) (err error) {
	var order model.Order
	//사용 시 orderId는 임시데이터로 넣을 것
	err = c.BodyParser(&order)
	if err != nil {
		c.Response().SetStatusCode(400)
		c.JSON(model.FailResponse{
			Message: "데이터를 읽는데 실패했습니다.",
			Error:   err,
		})
		return
	}

	//가상의 orderNumber를 만들어서 내려줌
	orderNum := common.GenerateRandomOrderNum()
	if err != nil {
		c.Response().SetStatusCode(400)
		c.JSON(err.Error())
		return
	}

	order.OrderId = orderNum

	c.Response().SetStatusCode(200)
	c.JSON(order)
	return
}
