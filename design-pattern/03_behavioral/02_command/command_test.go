package command

import (
	"strings"
	"testing"
)

func TestCommand(t *testing.T) {
	editor := NewEditor()
	invoker := NewCommandInvoker()

	// 测试插入命令
	insertCmd := NewInsertCommand(editor, "Hello")
	result := invoker.ExecuteCommand(insertCmd)
	if !strings.Contains(result, "插入文本: Hello") {
		t.Error("插入命令执行失败")
	}
	if editor.GetText() != "Hello" {
		t.Error("文本内容不正确")
	}

	// 测试第二次插入
	insertCmd2 := NewInsertCommand(editor, " World")
	invoker.ExecuteCommand(insertCmd2)
	if editor.GetText() != "Hello World" {
		t.Error("第二次插入失败")
	}

	// 测试删除命令
	deleteCmd := NewDeleteCommand(editor, 5)
	result = invoker.ExecuteCommand(deleteCmd)
	if !strings.Contains(result, "删除文本: World") {
		t.Error("删除命令执行失败")
	}
	if editor.GetText() != "Hello" {
		t.Error("删除后的文本内容不正确")
	}

	// 测试撤销
	result = invoker.Undo()
	if !strings.Contains(result, "撤销删除: World") {
		t.Error("撤销删除命令失败")
	}
	if editor.GetText() != "Hello World" {
		t.Error("撤销后的文本内容不正确")
	}

	// 测试多次撤销
	invoker.Undo() // 撤销第二次插入
	invoker.Undo() // 撤销第一次插入
	if editor.GetText() != "" {
		t.Error("多次撤销后文本应该为空")
	}

	// 测试撤销栈为空的情况
	result = invoker.Undo()
	if result != "没有可撤销的命令" {
		t.Error("空栈撤销处理不正确")
	}
}
