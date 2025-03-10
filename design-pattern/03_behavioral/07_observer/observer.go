package observer

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(message string)
	GetID() string
}

// Subject 主题接口
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll(message string)
}

// NewsAgency 新闻机构（具体主题）
type NewsAgency struct {
	observers map[string]Observer
}

func NewNewsAgency() *NewsAgency {
	return &NewsAgency{
		observers: make(map[string]Observer),
	}
}

func (n *NewsAgency) Register(observer Observer) {
	n.observers[observer.GetID()] = observer
}

func (n *NewsAgency) Deregister(observer Observer) {
	delete(n.observers, observer.GetID())
}

func (n *NewsAgency) NotifyAll(message string) {
	for _, observer := range n.observers {
		observer.Update(message)
	}
}

// NewsChannel 新闻频道（具体观察者）
type NewsChannel struct {
	id   string
	name string
	news []string
}

func NewNewsChannel(id, name string) *NewsChannel {
	return &NewsChannel{
		id:   id,
		name: name,
		news: make([]string, 0),
	}
}

func (n *NewsChannel) Update(message string) {
	n.news = append(n.news, message)
	fmt.Printf("频道 %s 收到新闻: %s\n", n.name, message)
}

func (n *NewsChannel) GetID() string {
	return n.id
}

func (n *NewsChannel) GetNews() []string {
	return n.news
}

// NewsReader 新闻读者（具体观察者）
type NewsReader struct {
	id       string
	name     string
	messages []string
}

func NewNewsReader(id, name string) *NewsReader {
	return &NewsReader{
		id:       id,
		name:     name,
		messages: make([]string, 0),
	}
}

func (n *NewsReader) Update(message string) {
	n.messages = append(n.messages, message)
	fmt.Printf("读者 %s 收到新闻: %s\n", n.name, message)
}

func (n *NewsReader) GetID() string {
	return n.id
}

func (n *NewsReader) GetMessages() []string {
	return n.messages
}

// EventType 事件类型
type EventType int

const (
	Breaking EventType = iota
	Normal
	Sports
)

// EventManager 事件管理器（支持按事件类型订阅）
type EventManager struct {
	listeners map[EventType]map[string]Observer
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[EventType]map[string]Observer),
	}
}

func (e *EventManager) Subscribe(eventType EventType, observer Observer) {
	if e.listeners[eventType] == nil {
		e.listeners[eventType] = make(map[string]Observer)
	}
	e.listeners[eventType][observer.GetID()] = observer
}

func (e *EventManager) Unsubscribe(eventType EventType, observer Observer) {
	if listeners, exists := e.listeners[eventType]; exists {
		delete(listeners, observer.GetID())
	}
}

func (e *EventManager) Notify(eventType EventType, message string) {
	if listeners, exists := e.listeners[eventType]; exists {
		for _, observer := range listeners {
			observer.Update(fmt.Sprintf("[%v] %s", eventType, message))
		}
	}
}
