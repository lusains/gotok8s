package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	// 创建聊天室（中介者）
	chatRoom := NewChatRoom()

	// 创建用户（同事）
	user1 := NewUser("张三")
	user2 := NewUser("李四")
	admin := NewAdminUser("管理员王五")

	// 注册用户到聊天室
	chatRoom.Register(user1)
	chatRoom.Register(user2)
	chatRoom.Register(admin)

	// 测试普通用户发送消息
	user1.Send("大家好！")

	// 测试管理员发送消息
	admin.Send("欢迎来到聊天室")

	// 测试管理员离线
	admin.SetOnline(false)
	admin.Send("这条消息不会被发送") // 应该看到离线提示

	// 测试管理员重新上线
	admin.SetOnline(true)
	admin.Send("我回来了")

	// 验证聊天室中的用户数量
	if len(chatRoom.colleagues) != 3 {
		t.Error("聊天室用户数量错误")
	}

	// 验证用户是否正确设置了中介者
	if user1.mediator != chatRoom {
		t.Error("用户1的中介者设置错误")
	}
	if user2.mediator != chatRoom {
		t.Error("用户2的中介者设置错误")
	}
	if admin.mediator != chatRoom {
		t.Error("管理员的中介者设置错误")
	}
}
