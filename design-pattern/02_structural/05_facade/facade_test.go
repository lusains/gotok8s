package facade

import "testing"

func TestFacade(t *testing.T) {
	// 测试基本外观
	computer := NewComputerFacade()
	computer.Start()

	// 测试扩展外观
	extendedComputer := NewExtendedComputerFacade()

	// 测试启动流程
	extendedComputer.StartComputer()

	// 测试关机流程
	extendedComputer.ShutDown()

	// 验证子系统是否正确初始化
	if extendedComputer.cpu == nil ||
		extendedComputer.memory == nil ||
		extendedComputer.hardDrive == nil ||
		extendedComputer.display == nil ||
		extendedComputer.os == nil {
		t.Error("子系统未正确初始化")
	}
}
