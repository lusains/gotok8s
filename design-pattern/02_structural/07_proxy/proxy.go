package proxy

import "time"

// Subject 抽象主题接口
type Subject interface {
	Request() string
}

// RealSubject 真实主题
type RealSubject struct{}

func (r *RealSubject) Request() string {
	// 模拟耗时操作
	time.Sleep(time.Second)
	return "RealSubject: 处理请求"
}

// Proxy 代理
type Proxy struct {
	realSubject *RealSubject
	cache       string
	accessCount int
}

func NewProxy() *Proxy {
	return &Proxy{
		realSubject: &RealSubject{},
	}
}

func (p *Proxy) Request() string {
	// 访问控制
	if p.accessCount >= 3 {
		return "Proxy: 访问次数超限"
	}

	// 缓存代理
	if p.cache != "" {
		return "Proxy: 返回缓存 - " + p.cache
	}

	// 记录请求
	println("Proxy: 记录请求时间:", time.Now().Format("15:04:05"))

	// 调用真实主题
	result := p.realSubject.Request()

	// 缓存结果
	p.cache = result
	p.accessCount++

	return result
}

// ProtectedProxy 保护代理
type ProtectedProxy struct {
	realSubject *RealSubject
	isAdmin     bool
}

func NewProtectedProxy(isAdmin bool) *ProtectedProxy {
	return &ProtectedProxy{
		realSubject: &RealSubject{},
		isAdmin:     isAdmin,
	}
}

func (p *ProtectedProxy) Request() string {
	if !p.isAdmin {
		return "ProtectedProxy: 没有访问权限"
	}
	return p.realSubject.Request()
}

// VirtualProxy 虚拟代理（延迟加载）
type VirtualProxy struct {
	realSubject *RealSubject
}

func NewVirtualProxy() *VirtualProxy {
	return &VirtualProxy{}
}

func (v *VirtualProxy) Request() string {
	// 延迟加载
	if v.realSubject == nil {
		println("VirtualProxy: 正在创建RealSubject...")
		v.realSubject = &RealSubject{}
	}
	return v.realSubject.Request()
}
