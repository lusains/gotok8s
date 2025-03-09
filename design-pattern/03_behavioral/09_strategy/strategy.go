package strategy

import "fmt"

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// AliPayStrategy 支付宝支付策略
type AliPayStrategy struct {
	userID string
}

func NewAliPayStrategy(userID string) *AliPayStrategy {
	return &AliPayStrategy{userID: userID}
}

func (a *AliPayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用支付宝支付 %.2f 元 (用户ID: %s)", amount, a.userID)
}

// WeChatPayStrategy 微信支付策略
type WeChatPayStrategy struct {
	openID string
}

func NewWeChatPayStrategy(openID string) *WeChatPayStrategy {
	return &WeChatPayStrategy{openID: openID}
}

func (w *WeChatPayStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用微信支付 %.2f 元 (OpenID: %s)", amount, w.openID)
}

// BankCardStrategy 银行卡支付策略
type BankCardStrategy struct {
	cardNumber string
	bankName   string
}

func NewBankCardStrategy(cardNumber, bankName string) *BankCardStrategy {
	return &BankCardStrategy{
		cardNumber: cardNumber,
		bankName:   bankName,
	}
}

func (b *BankCardStrategy) Pay(amount float64) string {
	return fmt.Sprintf("使用%s银行卡支付 %.2f 元 (卡号: %s)", b.bankName, amount, b.cardNumber)
}

// PaymentContext 支付上下文
type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) ExecutePayment(amount float64) string {
	if p.strategy == nil {
		return "未设置支付方式"
	}
	return p.strategy.Pay(amount)
}

// SortStrategy 排序策略接口
type SortStrategy interface {
	Sort(arr []int) []int
	GetName() string
}

// BubbleSortStrategy 冒泡排序策略
type BubbleSortStrategy struct{}

func (s *BubbleSortStrategy) Sort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func (s *BubbleSortStrategy) GetName() string {
	return "冒泡排序"
}

// QuickSortStrategy 快速排序策略
type QuickSortStrategy struct{}

func (s *QuickSortStrategy) Sort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	s.quickSort(result, 0, len(result)-1)
	return result
}

func (s *QuickSortStrategy) quickSort(arr []int, low, high int) {
	if low < high {
		pivot := s.partition(arr, low, high)
		s.quickSort(arr, low, pivot-1)
		s.quickSort(arr, pivot+1, high)
	}
}

func (s *QuickSortStrategy) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (s *QuickSortStrategy) GetName() string {
	return "快速排序"
}

// Sorter 排序器（上下文）
type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(arr []int) []int {
	if s.strategy == nil {
		return arr
	}
	fmt.Printf("使用%s算法\n", s.strategy.GetName())
	return s.strategy.Sort(arr)
}
