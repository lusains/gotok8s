package visitor

import "fmt"

// Shape 形状接口（元素接口）
type Shape interface {
	Accept(visitor ShapeVisitor)
	GetType() string
}

// ShapeVisitor 形状访问者接口
type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
	VisitTriangle(triangle *Triangle)
}

// Circle 圆形（具体元素）
type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

// Rectangle 矩形（具体元素）
type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) GetHeight() float64 {
	return r.height
}

// Triangle 三角形（具体元素）
type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(base, height float64) *Triangle {
	return &Triangle{
		base:   base,
		height: height,
	}
}

func (t *Triangle) Accept(visitor ShapeVisitor) {
	visitor.VisitTriangle(t)
}

func (t *Triangle) GetType() string {
	return "Triangle"
}

func (t *Triangle) GetBase() float64 {
	return t.base
}

func (t *Triangle) GetHeight() float64 {
	return t.height
}

// AreaCalculator 面积计算访问者（具体访问者）
type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) VisitCircle(circle *Circle) {
	a.area = 3.14 * circle.radius * circle.radius
}

func (a *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	a.area = rectangle.width * rectangle.height
}

func (a *AreaCalculator) VisitTriangle(triangle *Triangle) {
	a.area = 0.5 * triangle.base * triangle.height
}

func (a *AreaCalculator) GetArea() float64 {
	return a.area
}

// DrawVisitor 绘制访问者（具体访问者）
type DrawVisitor struct {
	output string
}

func (d *DrawVisitor) VisitCircle(circle *Circle) {
	d.output = fmt.Sprintf("Drawing Circle with radius: %.2f", circle.radius)
}

func (d *DrawVisitor) VisitRectangle(rectangle *Rectangle) {
	d.output = fmt.Sprintf("Drawing Rectangle with width: %.2f and height: %.2f",
		rectangle.width, rectangle.height)
}

func (d *DrawVisitor) VisitTriangle(triangle *Triangle) {
	d.output = fmt.Sprintf("Drawing Triangle with base: %.2f and height: %.2f",
		triangle.base, triangle.height)
}

func (d *DrawVisitor) GetOutput() string {
	return d.output
}

// XMLExportVisitor XML导出访问者（具体访问者）
type XMLExportVisitor struct {
	xml string
}

func (x *XMLExportVisitor) VisitCircle(circle *Circle) {
	x.xml = fmt.Sprintf("<circle><radius>%.2f</radius></circle>", circle.radius)
}

func (x *XMLExportVisitor) VisitRectangle(rectangle *Rectangle) {
	x.xml = fmt.Sprintf("<rectangle><width>%.2f</width><height>%.2f</height></rectangle>",
		rectangle.width, rectangle.height)
}

func (x *XMLExportVisitor) VisitTriangle(triangle *Triangle) {
	x.xml = fmt.Sprintf("<triangle><base>%.2f</base><height>%.2f</height></triangle>",
		triangle.base, triangle.height)
}

func (x *XMLExportVisitor) GetXML() string {
	return x.xml
}
