package composite

import (
	"strings"
	"testing"
)

func TestComposite(t *testing.T) {
	// 创建根目录
	root := NewComposite("root")

	// 创建子目录和文件
	folder1 := NewComposite("folder1")
	folder2 := NewComposite("folder2")
	file1 := NewLeaf("file1.txt")
	file2 := NewLeaf("file2.txt")
	file3 := NewLeaf("file3.txt")

	// 构建目录树
	root.Add(folder1)
	root.Add(folder2)
	folder1.Add(file1)
	folder1.Add(file2)
	folder2.Add(file3)

	// 测试目录结构
	result := root.Operation()
	expected := "Folder(root) [Folder(folder1) [File(file1.txt) File(file2.txt)] Folder(folder2) [File(file3.txt)]]"
	if result != expected {
		t.Errorf("目录结构不正确\n期望: %s\n实际: %s", expected, result)
	}

	// 测试删除操作
	folder1.Remove(file1)
	result = root.Operation()
	if strings.Contains(result, "file1.txt") {
		t.Error("删除文件失败")
	}

	// 测试获取子节点
	child := folder2.GetChild(0)
	if child.GetName() != "file3.txt" {
		t.Error("获取子节点失败")
	}

	// 测试叶子节点操作
	file1.Add(file2)    // 应该无效
	file1.Remove(file2) // 应该无效
	if file1.GetChild(0) != nil {
		t.Error("叶子节点不应该有子节点")
	}
}
