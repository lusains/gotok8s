package mediator

// Mediator 中介者接口
type Mediator interface {
	SendMessage(message string, colleague Colleague)
	Register(colleague Colleague)
}

// Colleague 同事接口
type Colleague interface {
	Send(message string)
	Receive(message string)
	SetMediator(mediator Mediator)
}

// ChatRoom 聊天室（具体中介者）
type ChatRoom struct {
	colleagues []Colleague
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		colleagues: make([]Colleague, 0),
	}
}

func (c *ChatRoom) SendMessage(message string, sender Colleague) {
	for _, colleague := range c.colleagues {
		// 发送消息给除了发送者之外的所有同事
		if colleague != sender {
			colleague.Receive(message)
		}
	}
}

func (c *ChatRoom) Register(colleague Colleague) {
	c.colleagues = append(c.colleagues, colleague)
	colleague.SetMediator(c)
}

// User 用户（具体同事）
type User struct {
	name     string
	mediator Mediator
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}

func (u *User) Send(message string) {
	println(u.name + " 发送消息: " + message)
	u.mediator.SendMessage(u.name+": "+message, u)
}

func (u *User) Receive(message string) {
	println(u.name + " 收到消息: " + message)
}

func (u *User) SetMediator(mediator Mediator) {
	u.mediator = mediator
}

// AdminUser 管理员用户（具体同事）
type AdminUser struct {
	User
	isOnline bool
}

func NewAdminUser(name string) *AdminUser {
	return &AdminUser{
		User:     User{name: name},
		isOnline: true,
	}
}

func (a *AdminUser) Send(message string) {
	if !a.isOnline {
		println(a.name + " 当前离线，无法发送消息")
		return
	}
	println("[管理员]" + a.name + " 发送消息: " + message)
	a.mediator.SendMessage("[管理员]"+a.name+": "+message, a)
}

func (a *AdminUser) SetOnline(online bool) {
	a.isOnline = online
	status := "上线"
	if !online {
		status = "离线"
	}
	if a.mediator != nil {
		a.mediator.SendMessage("[系统消息] 管理员"+a.name+status, a)
	}
}
