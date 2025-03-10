package proxy

import (
	"strings"
	"testing"
)

func TestProxy(t *testing.T) {
	// 测试基本代理
	proxy := NewProxy()

	// 第一次请求，应该访问真实对象
	result := proxy.Request()
	if result != "RealSubject: 处理请求" {
		t.Error("第一次请求应该访问真实对象")
	}

	// 第二次请求，应该使用缓存
	result = proxy.Request()
	if !strings.Contains(result, "返回缓存") {
		t.Error("第二次请求应该使用缓存")
	}

	// 测试访问次数限制
	proxy.Request()
	result = proxy.Request()
	if !strings.Contains(result, "访问次数超限") {
		t.Error("应该限制访问次数")
	}

	// 测试保护代理
	adminProxy := NewProtectedProxy(true)
	userProxy := NewProtectedProxy(false)

	if adminProxy.Request() != "RealSubject: 处理请求" {
		t.Error("管理员应该能够访问")
	}

	if !strings.Contains(userProxy.Request(), "没有访问权限") {
		t.Error("普通用户不应该能够访问")
	}

	// 测试虚拟代理
	virtualProxy := NewVirtualProxy()
	result = virtualProxy.Request()
	if result != "RealSubject: 处理请求" {
		t.Error("虚拟代理应该最终返回真实对象的结果")
	}
}
