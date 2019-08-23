package main

import (
	"testing"

	"github.com/JREAMLU/study/mocks"
)

func TestChargeCustomer(t *testing.T) {
	smsService := new(mocks.MessageService)

	// 然后我们将定义当 100 传递给 SendChargeNotification 时，需要返回什么
	// 在这里，我们希望它在成功发送通知后返回 true
	smsService.On("SendChargeNotification", 100).Return(nil)

	// 接下来，我们要定义要测试的服务
	myService := MyService{smsService}
	// 然后调用方法
	myService.ChargeCustomer(100)

	// 最后，我们验证 myService.ChargeCustomer 调用了我们模拟的 SendChargeNotification 方法
	smsService.AssertExpectations(t)
}
