package builder

import "testing"

func TestBuilder(t *testing.T) {
	// 创建建造者
	builder := NewGamingComputerBuilder()
	director := NewDirector(builder)

	// 构建游戏电脑
	gamingComputer := director.ConstructGamingComputer()
	if gamingComputer.CPU != "Intel i9" {
		t.Error("游戏电脑CPU配置错误")
	}
	if gamingComputer.GPU != "RTX 3080" {
		t.Error("游戏电脑GPU配置错误")
	}

	// 构建办公电脑
	officeComputer := director.ConstructOfficeComputer()
	if officeComputer.CPU != "Intel i5" {
		t.Error("办公电脑CPU配置错误")
	}
	if officeComputer.Memory != "16GB" {
		t.Error("办公电脑内存配置错误")
	}

	// 测试链式调用
	customComputer := builder.
		SetCPU("AMD Ryzen 7").
		SetMemory("64GB").
		SetDisk("4TB SSD").
		SetGPU("RTX 4090").
		Build()

	if customComputer.CPU != "AMD Ryzen 7" {
		t.Error("自定义电脑CPU配置错误")
	}
	if customComputer.Memory != "64GB" {
		t.Error("自定义电脑内存配置错误")
	}
}
