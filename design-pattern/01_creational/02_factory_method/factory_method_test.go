package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	// 创建文件日志记录器
	fileFactory := &FileLoggerFactory{}
	fileLogger := fileFactory.CreateLogger()
	fileLogger.Log("这是一条文件日志")

	// 创建控制台日志记录器
	consoleFactory := &ConsoleLoggerFactory{}
	consoleLogger := consoleFactory.CreateLogger()
	consoleLogger.Log("这是一条控制台日志")

	// 验证不同类型
	if _, ok := fileLogger.(*FileLogger); !ok {
		t.Error("fileLogger应该是FileLogger类型")
	}

	if _, ok := consoleLogger.(*ConsoleLogger); !ok {
		t.Error("consoleLogger应该是ConsoleLogger类型")
	}
}
