package app

import (
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func ApplyRoutes(r fiber.Router) {
	apis := r.Group("app", middleware.CheckStore())
	{
		// 	# 앱에서 타 앱에서 발행하는 이벤트를 구독하기 위한 리스너 API (앱 개발자가 URL Path 정의)
		apis.Post("/test", CreateOrderForTest) // -> 해당 api에 이벤트를 등록 (임시 트리거)
	}
}

// 알림 사료, 간식 구매 시기를 저장하고 평균 교체 주기를 확인  -> 구매 시기가 되면 알림
//초기에는 임시로 데이터 설정하고 평균 구함

// - 이벤트 발행, 구독 가능
//  - 발행이벤트 : 다음 주기를 알려줌 < 구독자들은 구매 예상시기를 알 수 있음
//  - 구독이벤트 : 구매현황을 구독가능 -> 구매내역 받아서 DB에 추가하는 기능

// - 앱 호출 대응 가능
// 	- 화면 확장 가능(optional) - 아무거나 파일 올리기

// - CRUD 가능
// 	- 내새꾸이름, 생일, 몸무개 입력  - 업데이트
// 	- 사료, 간식 구매시기 입력 - 수정
// 	- 조회
// 	- 삭제

// - 토큰 저장 기능
// - 앱의 파일 업로드 시 웹훅 받을 수 있는 API

// 	# 앱에서 타 앱에서 발행하는 이벤트를 구독하기 위한 리스너 API (앱 개발자가 URL Path 정의)
// 	apis.Post("/order", CreateOrder) //-> 구매내역을 받아서 추가 (test)api를 구독 (임시로 진짜 결제 api로 대체 가능) 얘가 1번 리스너가 될 것
// 	apis.Put("/order", UpdateOrder)
// 	apis.Get("/order", GetOrderByOrderId)
// 	apis.Delete("/order", OrderId)
// 	apis.Get("/alram", RemindShoppingDayByUserId)

// }
