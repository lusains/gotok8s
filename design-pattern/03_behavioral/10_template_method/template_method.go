package template_method

import "fmt"

// DataMiner 数据挖掘抽象类
type DataMiner interface {
	Mine() // 模板方法
	OpenFile() string
	ExtractData() []string
	ParseData([]string) []interface{}
	AnalyzeData([]interface{}) string
	SendReport(string)
	Hook() bool // 钩子方法
}

// BaseDataMiner 基础数据挖掘类
type BaseDataMiner struct {
	implementation DataMiner
}

// Mine 模板方法，定义了算法骨架
func (b *BaseDataMiner) Mine() {
	if b.implementation == nil {
		return
	}

	// 1. 打开文件
	filename := b.implementation.OpenFile()
	fmt.Printf("打开文件: %s\n", filename)

	// 2. 提取数据
	rawData := b.implementation.ExtractData()
	fmt.Printf("提取原始数据: %v\n", rawData)

	// 3. 解析数据
	parsedData := b.implementation.ParseData(rawData)
	fmt.Printf("解析后的数据: %v\n", parsedData)

	// 4. 分析数据
	result := b.implementation.AnalyzeData(parsedData)
	fmt.Printf("分析结果: %s\n", result)

	// 钩子方法：决定是否发送报告
	if b.implementation.Hook() {
		// 5. 发送报告
		b.implementation.SendReport(result)
	}
}

// PDFDataMiner PDF数据挖掘器
type PDFDataMiner struct {
	BaseDataMiner
	filename string
}

func NewPDFDataMiner(filename string) *PDFDataMiner {
	miner := &PDFDataMiner{
		filename: filename,
	}
	miner.implementation = miner
	return miner
}

func (p *PDFDataMiner) OpenFile() string {
	return "PDF文件: " + p.filename
}

func (p *PDFDataMiner) ExtractData() []string {
	return []string{"PDF数据1", "PDF数据2", "PDF数据3"}
}

func (p *PDFDataMiner) ParseData(data []string) []interface{} {
	var result []interface{}
	for _, d := range data {
		result = append(result, "已解析:"+d)
	}
	return result
}

func (p *PDFDataMiner) AnalyzeData(data []interface{}) string {
	return fmt.Sprintf("PDF分析结果: 共%d条数据", len(data))
}

func (p *PDFDataMiner) SendReport(result string) {
	fmt.Printf("发送PDF报告: %s\n", result)
}

func (p *PDFDataMiner) Hook() bool {
	return true // PDF文件总是发送报告
}

// CSVDataMiner CSV数据挖掘器
type CSVDataMiner struct {
	BaseDataMiner
	filename   string
	sendReport bool
}

func NewCSVDataMiner(filename string, sendReport bool) *CSVDataMiner {
	miner := &CSVDataMiner{
		filename:   filename,
		sendReport: sendReport,
	}
	miner.implementation = miner
	return miner
}

func (c *CSVDataMiner) OpenFile() string {
	return "CSV文件: " + c.filename
}

func (c *CSVDataMiner) ExtractData() []string {
	return []string{"CSV行1", "CSV行2", "CSV行3", "CSV行4"}
}

func (c *CSVDataMiner) ParseData(data []string) []interface{} {
	var result []interface{}
	for i, d := range data {
		result = append(result, fmt.Sprintf("第%d行:%s", i+1, d))
	}
	return result
}

func (c *CSVDataMiner) AnalyzeData(data []interface{}) string {
	return fmt.Sprintf("CSV分析结果: 包含%d行数据", len(data))
}

func (c *CSVDataMiner) SendReport(result string) {
	fmt.Printf("发送CSV报告: %s\n", result)
}

func (c *CSVDataMiner) Hook() bool {
	return c.sendReport // 根据配置决定是否发送报告
}
