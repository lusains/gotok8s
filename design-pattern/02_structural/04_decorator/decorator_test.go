package decorator

import "testing"

func TestDecorator(t *testing.T) {
	// 创建一个简单的数据源
	source := NewFileDataSource("test.txt")

	// 测试原始数据源
	source.WriteData("Hello World")
	if data := source.ReadData(); data != "Hello World" {
		t.Error("原始数据源读写失败")
	}

	// 测试加密装饰器
	encrypted := NewEncryptionDecorator(source)
	encrypted.WriteData("Hello World")
	if data := encrypted.ReadData(); data != "Hello World" {
		t.Errorf("加密装饰器失败，期望 'Hello World'，得到 '%s'", data)
	}

	// 测试压缩装饰器
	compressed := NewCompressionDecorator(source)
	compressed.WriteData("Hello World")
	if data := compressed.ReadData(); data != "Hello World" {
		t.Errorf("压缩装饰器失败，期望 'Hello World'，得到 '%s'", data)
	}

	// 测试组合装饰器（压缩+加密）
	encryptedAndCompressed := NewEncryptionDecorator(NewCompressionDecorator(source))
	encryptedAndCompressed.WriteData("Hello World")
	if data := encryptedAndCompressed.ReadData(); data != "Hello World" {
		t.Errorf("组合装饰器失败，期望 'Hello World'，得到 '%s'", data)
	}
}
