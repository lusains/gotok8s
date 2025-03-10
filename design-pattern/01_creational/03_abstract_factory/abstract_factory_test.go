package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	// 测试暗色主题
	darkFactory := &DarkFactory{}
	darkButton := darkFactory.CreateButton()
	darkInput := darkFactory.CreateInput()

	if darkButton.Render() != "渲染暗色按钮" {
		t.Error("暗色按钮渲染错误")
	}
	darkButton.OnClick()
	darkInput.OnInput("测试暗色输入")

	// 测试亮色主题
	lightFactory := &LightFactory{}
	lightButton := lightFactory.CreateButton()
	lightInput := lightFactory.CreateInput()

	if lightButton.Render() != "渲染亮色按钮" {
		t.Error("亮色按钮渲染错误")
	}
	lightButton.OnClick()
	lightInput.OnInput("测试亮色输入")

	// 验证类型
	if _, ok := darkButton.(*DarkButton); !ok {
		t.Error("darkButton类型错误")
	}
	if _, ok := lightInput.(*LightInput); !ok {
		t.Error("lightInput类型错误")
	}
}
