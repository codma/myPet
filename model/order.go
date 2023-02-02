package model

type Order struct {
	OrderId     string `json:"order_id" validate:"required"`  //주문시 발생하는 orderId
	UserName    string `json:"user_name" validate:"required"` //펫네임
	Category    string `json:"categody" validate:"required"`  //주식 or 간식
	ProductName string `json:"product_name" `                 //참고용
}
