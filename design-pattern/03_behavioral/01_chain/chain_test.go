package chain

import (
	"strings"
	"testing"
)

func TestChain(t *testing.T) {
	// 创建处理链
	auth := &AuthHandler{}
	validation := &ValidationHandler{}
	log := &LogHandler{}
	business := &BusinessHandler{}

	auth.SetNext(validation).SetNext(log).SetNext(business)

	// 测试有效请求
	result := auth.Handle("admin")
	if !strings.Contains(result, "认证成功") ||
		!strings.Contains(result, "验证通过") ||
		!strings.Contains(result, "记录请求") ||
		!strings.Contains(result, "处理业务逻辑") {
		t.Error("有效请求应该通过所有处理者")
	}

	// 测试认证失败
	result = auth.Handle("user")
	if !strings.Contains(result, "认证失败") {
		t.Error("非admin用户应该认证失败")
	}

	// 测试验证失败
	result = auth.Handle("ad")
	if !strings.Contains(result, "验证失败") {
		t.Error("短请求应该验证失败")
	}

	// 测试空链
	emptyHandler := &BaseHandler{}
	if result := emptyHandler.Handle("test"); result != "" {
		t.Error("空链应该返回空字符串")
	}
}
