package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {
	editor := NewEditor()
	history := NewHistory()

	// 测试输入文本并保存状态
	editor.Type("Hello")
	history.Push(editor.Save())
	if editor.GetContent() != "Hello" {
		t.Error("文本输入错误")
	}

	editor.Type(" World")
	history.Push(editor.Save())
	if editor.GetContent() != "Hello World" {
		t.Error("文本输入错误")
	}

	editor.Type("!")
	history.Push(editor.Save())
	if editor.GetContent() != "Hello World!" {
		t.Error("文本输入错误")
	}

	// 测试撤销
	editor.Restore(history.Undo())
	if editor.GetContent() != "Hello World" {
		t.Error("撤销操作失败")
	}

	editor.Restore(history.Undo())
	if editor.GetContent() != "Hello" {
		t.Error("撤销操作失败")
	}

	// 测试重做
	editor.Restore(history.Redo())
	if editor.GetContent() != "Hello World" {
		t.Error("重做操作失败")
	}

	// 测试在中间状态编辑
	editor.Type(" Again")
	history.Push(editor.Save())
	if editor.GetContent() != "Hello World Again" {
		t.Error("中间状态编辑失败")
	}

	// 验证历史记录
	histories := history.GetHistory()
	if len(histories) != 4 {
		t.Error("历史记录数量错误")
	}

	// 验证时间戳
	memento := editor.Save()
	if memento.GetTimestamp().IsZero() {
		t.Error("时间戳未正确设置")
	}
}
