package chain

// Handler 处理者接口
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// BaseHandler 基础处理者
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) string {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return ""
}

// AuthHandler 认证处理者
type AuthHandler struct {
	BaseHandler
}

func (a *AuthHandler) Handle(request string) string {
	if request == "admin" {
		return "AuthHandler: 认证成功\n" + a.BaseHandler.Handle(request)
	}
	return "AuthHandler: 认证失败"
}

// ValidationHandler 验证处理者
type ValidationHandler struct {
	BaseHandler
}

func (v *ValidationHandler) Handle(request string) string {
	if len(request) > 3 {
		return "ValidationHandler: 验证通过\n" + v.BaseHandler.Handle(request)
	}
	return "ValidationHandler: 验证失败，请求长度必须大于3"
}

// LogHandler 日志处理者
type LogHandler struct {
	BaseHandler
}

func (l *LogHandler) Handle(request string) string {
	return "LogHandler: 记录请求 - " + request + "\n" + l.BaseHandler.Handle(request)
}

// BusinessHandler 业务处理者
type BusinessHandler struct {
	BaseHandler
}

func (b *BusinessHandler) Handle(request string) string {
	return "BusinessHandler: 处理业务逻辑 - " + request
}
