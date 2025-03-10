package template_method

import (
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	// 测试PDF数据挖掘器
	pdfMiner := NewPDFDataMiner("sample.pdf")
	pdfMiner.Mine()

	// 测试CSV数据挖掘器（发送报告）
	csvMiner := NewCSVDataMiner("data.csv", true)
	csvMiner.Mine()

	// 测试CSV数据挖掘器（不发送报告）
	csvMinerNoReport := NewCSVDataMiner("data.csv", false)
	csvMinerNoReport.Mine()

	// 测试基础数据挖掘器（应该不执行任何操作）
	var baseMiner BaseDataMiner
	baseMiner.Mine()
}

func TestPDFDataMiner(t *testing.T) {
	miner := NewPDFDataMiner("test.pdf")

	// 测试各个步骤
	if filename := miner.OpenFile(); filename != "PDF文件: test.pdf" {
		t.Error("PDF文件名错误")
	}

	if data := miner.ExtractData(); len(data) != 3 {
		t.Error("PDF数据提取数量错误")
	}

	rawData := []string{"测试数据"}
	if parsedData := miner.ParseData(rawData); len(parsedData) != 1 {
		t.Error("PDF数据解析错误")
	}

	if result := miner.AnalyzeData([]interface{}{1, 2, 3}); result != "PDF分析结果: 共3条数据" {
		t.Error("PDF数据分析结果错误")
	}

	if !miner.Hook() {
		t.Error("PDF钩子方法应该返回true")
	}
}

func TestCSVDataMiner(t *testing.T) {
	miner := NewCSVDataMiner("test.csv", true)

	// 测试各个步骤
	if filename := miner.OpenFile(); filename != "CSV文件: test.csv" {
		t.Error("CSV文件名错误")
	}

	if data := miner.ExtractData(); len(data) != 4 {
		t.Error("CSV数据提取数量错误")
	}

	rawData := []string{"测试数据"}
	if parsedData := miner.ParseData(rawData); len(parsedData) != 1 {
		t.Error("CSV数据解析错误")
	}

	if result := miner.AnalyzeData([]interface{}{1, 2, 3, 4}); result != "CSV分析结果: 包含4行数据" {
		t.Error("CSV数据分析结果错误")
	}

	// 测试钩子方法
	if !miner.Hook() {
		t.Error("CSV钩子方法配置为true时应该返回true")
	}

	minerNoReport := NewCSVDataMiner("test.csv", false)
	if minerNoReport.Hook() {
		t.Error("CSV钩子方法配置为false时应该返回false")
	}
}
