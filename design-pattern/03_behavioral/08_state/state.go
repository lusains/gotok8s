package state

import "fmt"

// State 状态接口
type State interface {
	Handle(order *Order)
	String() string
}

// Order 订单（上下文）
type Order struct {
	state State
}

func NewOrder() *Order {
	order := &Order{}
	order.state = &NewState{} // 初始状态为新建
	return order
}

func (o *Order) SetState(state State) {
	o.state = state
}

func (o *Order) GetState() State {
	return o.state
}

func (o *Order) Process() {
	o.state.Handle(o)
}

// NewState 新建状态
type NewState struct{}

func (s *NewState) Handle(order *Order) {
	fmt.Println("订单已创建，等待付款...")
	order.SetState(&PaidState{})
}

func (s *NewState) String() string {
	return "新建"
}

// PaidState 已付款状态
type PaidState struct{}

func (s *PaidState) Handle(order *Order) {
	fmt.Println("订单已付款，准备发货...")
	order.SetState(&ShippedState{})
}

func (s *PaidState) String() string {
	return "已付款"
}

// ShippedState 已发货状态
type ShippedState struct{}

func (s *ShippedState) Handle(order *Order) {
	fmt.Println("订单已发货，等待签收...")
	order.SetState(&ReceivedState{})
}

func (s *ShippedState) String() string {
	return "已发货"
}

// ReceivedState 已签收状态
type ReceivedState struct{}

func (s *ReceivedState) Handle(order *Order) {
	fmt.Println("订单已签收，交易完成")
	order.SetState(&CompletedState{})
}

func (s *ReceivedState) String() string {
	return "已签收"
}

// CompletedState 已完成状态
type CompletedState struct{}

func (s *CompletedState) Handle(order *Order) {
	fmt.Println("订单已完成，无需进一步处理")
}

func (s *CompletedState) String() string {
	return "已完成"
}

// CancelledState 已取消状态
type CancelledState struct{}

func (s *CancelledState) Handle(order *Order) {
	fmt.Println("订单已取消，无法进一步处理")
}

func (s *CancelledState) String() string {
	return "已取消"
}

// OrderManager 订单管理器
type OrderManager struct {
	order *Order
}

func NewOrderManager() *OrderManager {
	return &OrderManager{
		order: NewOrder(),
	}
}

func (m *OrderManager) ProcessOrder() {
	m.order.Process()
}

func (m *OrderManager) GetCurrentState() string {
	return m.order.GetState().String()
}

func (m *OrderManager) CancelOrder() {
	fmt.Println("订单取消中...")
	m.order.SetState(&CancelledState{})
}
