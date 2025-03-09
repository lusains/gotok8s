package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	// 获取单例实例
	instance1 := GetInstance()
	instance2 := GetInstance()

	// 验证两个实例是否是同一个对象
	if instance1 != instance2 {
		t.Error("实例应该相同")
	}

	// 测试数据操作
	instance1.SetData("新数据")
	if instance2.GetData() != "新数据" {
		t.Error("数据不一致")
	}
}
