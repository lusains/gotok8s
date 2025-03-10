package observer

import (
	"strings"
	"testing"
)

func TestObserver(t *testing.T) {
	// 创建新闻机构
	newsAgency := NewNewsAgency()

	// 创建观察者
	channel1 := NewNewsChannel("ch1", "CCTV1")
	channel2 := NewNewsChannel("ch2", "CCTV2")
	reader1 := NewNewsReader("r1", "张三")

	// 注册观察者
	newsAgency.Register(channel1)
	newsAgency.Register(channel2)
	newsAgency.Register(reader1)

	// 发送新闻
	newsAgency.NotifyAll("重大新闻：观察者模式发布！")

	// 验证观察者是否收到消息
	if len(channel1.GetNews()) != 1 {
		t.Error("频道1未收到消息")
	}
	if len(channel2.GetNews()) != 1 {
		t.Error("频道2未收到消息")
	}
	if len(reader1.GetMessages()) != 1 {
		t.Error("读者1未收到消息")
	}

	// 测试取消注册
	newsAgency.Deregister(channel2)
	newsAgency.NotifyAll("第二条新闻")

	if len(channel2.GetNews()) != 1 {
		t.Error("已取消注册的频道2不应该收到新消息")
	}
	if len(channel1.GetNews()) != 2 {
		t.Error("频道1应该收到两条消息")
	}

	// 测试事件管理器
	eventManager := NewEventManager()
	eventManager.Subscribe(Breaking, channel1)
	eventManager.Subscribe(Sports, channel2)
	eventManager.Subscribe(Breaking, reader1)

	// 发送不同类型的事件
	eventManager.Notify(Breaking, "突发新闻！")
	eventManager.Notify(Sports, "体育新闻")

	// 验证消息接收
	found := false
	for _, news := range channel1.GetNews() {
		if strings.Contains(news, "突发新闻") {
			found = true
			break
		}
	}
	if !found {
		t.Error("频道1未收到突发新闻")
	}

	// 测试取消订阅
	eventManager.Unsubscribe(Breaking, channel1)
	eventManager.Notify(Breaking, "另一条突发新闻")
	if len(channel1.GetNews()) != 3 {
		t.Error("取消订阅后不应该收到新消息")
	}
}
