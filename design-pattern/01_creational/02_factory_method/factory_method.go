package factory_method

// Logger 定义日志记录器接口
type Logger interface {
	Log(message string)
}

// LoggerFactory 定义日志工厂接口
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLogger 文件日志记录器
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	// 实际场景中会写入文件
	println("File Logger:", message)
}

// ConsoleLogger 控制台日志记录器
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	println("Console Logger:", message)
}

// FileLoggerFactory 文件日志工厂
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory 控制台日志工厂
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}
