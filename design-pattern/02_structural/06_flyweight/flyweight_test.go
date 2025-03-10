package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	forest := NewForest()

	// 种植不同类型的树
	forest.PlantTree(1, 1, "杉树", "绿色", "粗糙")
	forest.PlantTree(5, 5, "杉树", "绿色", "粗糙") // 应该复用已有的树类型
	forest.PlantTree(10, 10, "松树", "深绿", "光滑")

	// 验证树类型是否被正确共享
	treeTypeCount := len(forest.factory.treeTypes)
	if treeTypeCount != 2 {
		t.Errorf("期望树类型数量为2，实际为%d", treeTypeCount)
	}

	// 验证树的数量
	treeCount := len(forest.trees)
	if treeCount != 3 {
		t.Errorf("期望树的数量为3，实际为%d", treeCount)
	}

	// 测试绘制
	forest.Draw()

	// 测试工厂的复用能力
	factory := NewTreeFactory()
	type1 := factory.GetTreeType("杉树", "绿色", "粗糙")
	type2 := factory.GetTreeType("杉树", "绿色", "粗糙")
	if type1 != type2 {
		t.Error("相同参数应该返回同一个树类型实例")
	}
}
