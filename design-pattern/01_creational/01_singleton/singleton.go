package singleton

import "sync"

// Singleton 定义单例结构体
type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

// GetInstance 获取单例实例的方法
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			data: "这是单例中的数据",
		}
	})
	return instance
}

// GetData 获取数据的方法
func (s *Singleton) GetData() string {
	return s.data
}

// SetData 设置数据的方法
func (s *Singleton) SetData(data string) {
	s.data = data
}
