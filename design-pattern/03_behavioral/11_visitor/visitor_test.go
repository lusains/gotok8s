package visitor

import (
	"strings"
	"testing"
)

func TestVisitor(t *testing.T) {
	// 创建形状
	circle := NewCircle(5)
	rectangle := NewRectangle(4, 6)
	triangle := NewTriangle(3, 4)

	// 测试面积计算访问者
	areaCalculator := &AreaCalculator{}

	circle.Accept(areaCalculator)
	if area := areaCalculator.GetArea(); area != 78.5 {
		t.Errorf("圆形面积计算错误，期望：78.5，实际：%.2f", area)
	}

	rectangle.Accept(areaCalculator)
	if area := areaCalculator.GetArea(); area != 24 {
		t.Errorf("矩形面积计算错误，期望：24，实际：%.2f", area)
	}

	triangle.Accept(areaCalculator)
	if area := areaCalculator.GetArea(); area != 6 {
		t.Errorf("三角形面积计算错误，期望：6，实际：%.2f", area)
	}

	// 测试绘制访问者
	drawVisitor := &DrawVisitor{}

	circle.Accept(drawVisitor)
	if output := drawVisitor.GetOutput(); !strings.Contains(output, "Circle") {
		t.Error("圆形绘制输出错误")
	}

	rectangle.Accept(drawVisitor)
	if output := drawVisitor.GetOutput(); !strings.Contains(output, "Rectangle") {
		t.Error("矩形绘制输出错误")
	}

	triangle.Accept(drawVisitor)
	if output := drawVisitor.GetOutput(); !strings.Contains(output, "Triangle") {
		t.Error("三角形绘制输出错误")
	}

	// 测试XML导出访问者
	xmlExporter := &XMLExportVisitor{}

	circle.Accept(xmlExporter)
	if xml := xmlExporter.GetXML(); !strings.Contains(xml, "<circle>") {
		t.Error("圆形XML导出错误")
	}

	rectangle.Accept(xmlExporter)
	if xml := xmlExporter.GetXML(); !strings.Contains(xml, "<rectangle>") {
		t.Error("矩形XML导出错误")
	}

	triangle.Accept(xmlExporter)
	if xml := xmlExporter.GetXML(); !strings.Contains(xml, "<triangle>") {
		t.Error("三角形XML导出错误")
	}

	// 测试元素类型
	if circle.GetType() != "Circle" {
		t.Error("圆形类型错误")
	}
	if rectangle.GetType() != "Rectangle" {
		t.Error("矩形类型错误")
	}
	if triangle.GetType() != "Triangle" {
		t.Error("三角形类型错误")
	}
}
