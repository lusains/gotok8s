package decorator

// DataSource 数据源接口（抽象组件）
type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// FileDataSource 文件数据源（具体组件）
type FileDataSource struct {
	filename string
	data     string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{
		filename: filename,
	}
}

func (f *FileDataSource) WriteData(data string) {
	// 实际应用中会写入文件，这里简化为保存到内存
	f.data = data
	println("写入文件数据:", data)
}

func (f *FileDataSource) ReadData() string {
	println("读取文件数据:", f.data)
	return f.data
}

// DataSourceDecorator 数据源装饰器（抽象装饰器）
type DataSourceDecorator struct {
	wrappee DataSource
}

func NewDataSourceDecorator(source DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{
		wrappee: source,
	}
}

func (d *DataSourceDecorator) WriteData(data string) {
	d.wrappee.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() string {
	return d.wrappee.ReadData()
}

// EncryptionDecorator 加密装饰器（具体装饰器）
type EncryptionDecorator struct {
	*DataSourceDecorator
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{
		DataSourceDecorator: NewDataSourceDecorator(source),
	}
}

func (e *EncryptionDecorator) WriteData(data string) {
	// 加密数据（这里用简单的反转字符串模拟加密）
	encrypted := reverse(data)
	e.wrappee.WriteData(encrypted)
}

func (e *EncryptionDecorator) ReadData() string {
	// 解密数据
	data := e.wrappee.ReadData()
	return reverse(data)
}

// CompressionDecorator 压缩装饰器（具体装饰器）
type CompressionDecorator struct {
	*DataSourceDecorator
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{
		DataSourceDecorator: NewDataSourceDecorator(source),
	}
}

func (c *CompressionDecorator) WriteData(data string) {
	// 压缩数据（这里用简单的添加标记模拟压缩）
	compressed := "COMPRESSED:" + data
	c.wrappee.WriteData(compressed)
}

func (c *CompressionDecorator) ReadData() string {
	// 解压数据
	data := c.wrappee.ReadData()
	return data[11:] // 移除 "COMPRESSED:" 前缀
}

// 辅助函数：反转字符串
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
