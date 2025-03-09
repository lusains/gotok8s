package prototype

import "encoding/json"

// Prototype 原型接口
type Prototype interface {
	Clone() Prototype
	DeepClone() Prototype
}

// Document 文档结构体
type Document struct {
	Title    string
	Content  string
	Author   *Author
	Metadata map[string]string
}

// Author 作者结构体
type Author struct {
	Name  string
	Email string
}

// Clone 浅拷贝实现
func (d *Document) Clone() Prototype {
	// 浅拷贝会复制指针，新旧对象共享相同的引用类型数据
	docClone := *d
	return &docClone
}

// DeepClone 深拷贝实现
func (d *Document) DeepClone() Prototype {
	// 使用JSON序列化和反序列化实现深拷贝
	bytes, _ := json.Marshal(d)
	var docClone Document
	json.Unmarshal(bytes, &docClone)
	return &docClone
}

// PrototypeManager 原型管理器
type PrototypeManager struct {
	prototypes map[string]Prototype
}

// NewPrototypeManager 创建原型管理器
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Prototype),
	}
}

// Register 注册原型
func (p *PrototypeManager) Register(name string, prototype Prototype) {
	p.prototypes[name] = prototype
}

// Get 获取原型
func (p *PrototypeManager) Get(name string) Prototype {
	if prototype, ok := p.prototypes[name]; ok {
		return prototype.DeepClone()
	}
	return nil
}
