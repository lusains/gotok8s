package state

import (
	"testing"
)

func TestState(t *testing.T) {
	manager := NewOrderManager()

	// 测试初始状态
	if manager.GetCurrentState() != "新建" {
		t.Error("初始状态应该是'新建'")
	}

	// 测试订单处理流程
	manager.ProcessOrder() // 新建 -> 已付款
	if manager.GetCurrentState() != "已付款" {
		t.Error("第一次处理后状态应该是'已付款'")
	}

	manager.ProcessOrder() // 已付款 -> 已发货
	if manager.GetCurrentState() != "已发货" {
		t.Error("第二次处理后状态应该是'已发货'")
	}

	manager.ProcessOrder() // 已发货 -> 已签收
	if manager.GetCurrentState() != "已签收" {
		t.Error("第三次处理后状态应该是'已签收'")
	}

	manager.ProcessOrder() // 已签收 -> 已完成
	if manager.GetCurrentState() != "已完成" {
		t.Error("第四次处理后状态应该是'已完成'")
	}

	// 测试完成状态的处理
	manager.ProcessOrder() // 已完成状态下继续处理
	if manager.GetCurrentState() != "已完成" {
		t.Error("完成状态应该保持不变")
	}

	// 测试订单取消
	manager = NewOrderManager() // 创建新订单
	manager.CancelOrder()
	if manager.GetCurrentState() != "已取消" {
		t.Error("取消后状态应该是'已取消'")
	}

	// 测试取消状态的处理
	manager.ProcessOrder() // 已取消状态下继续处理
	if manager.GetCurrentState() != "已取消" {
		t.Error("取消状态应该保持不变")
	}
}
