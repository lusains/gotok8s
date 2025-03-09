package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	// 创建原始文档
	originalDoc := &Document{
		Title:   "原型模式",
		Content: "这是关于原型模式的文档",
		Author: &Author{
			Name:  "张三",
			Email: "zhangsan@example.com",
		},
		Metadata: map[string]string{
			"version": "1.0",
			"status":  "draft",
		},
	}

	// 测试浅拷贝
	shallowCopy := originalDoc.Clone().(*Document)
	if shallowCopy.Author != originalDoc.Author {
		t.Error("浅拷贝应该共享Author引用")
	}

	// 测试深拷贝
	deepCopy := originalDoc.DeepClone().(*Document)
	if deepCopy.Author == originalDoc.Author {
		t.Error("深拷贝不应该共享Author引用")
	}

	// 修改原始文档，验证深拷贝的独立性
	originalDoc.Author.Name = "李四"
	if deepCopy.Author.Name == "李四" {
		t.Error("深拷贝的Author不应受原始对象修改影响")
	}

	// 测试原型管理器
	manager := NewPrototypeManager()
	manager.Register("doc", originalDoc)

	// 从管理器获取新实例
	newDoc := manager.Get("doc").(*Document)
	if newDoc == originalDoc {
		t.Error("管理器应该返回一个新的实例")
	}
}
