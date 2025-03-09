package strategy

import (
	"reflect"
	"testing"
)

func TestPaymentStrategy(t *testing.T) {
	// 创建支付上下文
	paymentContext := NewPaymentContext(nil)

	// 测试未设置策略的情况
	result := paymentContext.ExecutePayment(100)
	if result != "未设置支付方式" {
		t.Error("未设置策略时应该返回错误信息")
	}

	// 测试支付宝支付
	alipay := NewAliPayStrategy("12345")
	paymentContext.SetStrategy(alipay)
	result = paymentContext.ExecutePayment(100)
	expected := "使用支付宝支付 100.00 元 (用户ID: 12345)"
	if result != expected {
		t.Errorf("支付宝支付结果错误，期望：%s，实际：%s", expected, result)
	}

	// 测试微信支付
	wechat := NewWeChatPayStrategy("wx123")
	paymentContext.SetStrategy(wechat)
	result = paymentContext.ExecutePayment(200)
	expected = "使用微信支付 200.00 元 (OpenID: wx123)"
	if result != expected {
		t.Errorf("微信支付结果错误，期望：%s，实际：%s", expected, result)
	}

	// 测试银行卡支付
	bankCard := NewBankCardStrategy("6222021234", "工商")
	paymentContext.SetStrategy(bankCard)
	result = paymentContext.ExecutePayment(300)
	expected = "使用工商银行卡支付 300.00 元 (卡号: 6222021234)"
	if result != expected {
		t.Errorf("银行卡支付结果错误，期望：%s，实际：%s", expected, result)
	}
}

func TestSortStrategy(t *testing.T) {
	// 准备测试数据
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}

	// 测试冒泡排序
	sorter := NewSorter(&BubbleSortStrategy{})
	result := sorter.Sort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("冒泡排序结果错误，期望：%v，实际：%v", expected, result)
	}

	// 测试快速排序
	sorter.SetStrategy(&QuickSortStrategy{})
	result = sorter.Sort(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("快速排序结果错误，期望：%v，实际：%v", expected, result)
	}

	// 测试原数组不被修改
	original := []int{64, 34, 25, 12, 22, 11, 90}
	if !reflect.DeepEqual(arr, original) {
		t.Error("原数组不应被修改")
	}

	// 测试空策略
	sorter.SetStrategy(nil)
	result = sorter.Sort(arr)
	if !reflect.DeepEqual(result, arr) {
		t.Error("空策略应该返回原数组")
	}
}
