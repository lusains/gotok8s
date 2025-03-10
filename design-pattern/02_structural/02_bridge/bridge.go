package bridge

// DrawAPI 绘图API接口（实现类接口）
type DrawAPI interface {
	DrawCircle(x, y, radius int)
	DrawRectangle(x, y, width, height int)
}

// RedCircle 红色绘图API实现
type RedCircle struct{}

func (r *RedCircle) DrawCircle(x, y, radius int) {
	println("Drawing Circle[ color: red, x:", x, ", y:", y, ", radius:", radius, "]")
}

func (r *RedCircle) DrawRectangle(x, y, width, height int) {
	println("Drawing Rectangle[ color: red, x:", x, ", y:", y, ", width:", width, ", height:", height, "]")
}

// GreenCircle 绿色绘图API实现
type GreenCircle struct{}

func (g *GreenCircle) DrawCircle(x, y, radius int) {
	println("Drawing Circle[ color: green, x:", x, ", y:", y, ", radius:", radius, "]")
}

func (g *GreenCircle) DrawRectangle(x, y, width, height int) {
	println("Drawing Rectangle[ color: green, x:", x, ", y:", y, ", width:", width, ", height:", height, "]")
}

// Shape 形状抽象类
type Shape struct {
	drawAPI DrawAPI
}

// NewShape 创建形状
func NewShape(drawAPI DrawAPI) *Shape {
	return &Shape{drawAPI: drawAPI}
}

// Circle 圆形
type Circle struct {
	*Shape
	x, y, radius int
}

// NewCircle 创建圆形
func NewCircle(x, y, radius int, drawAPI DrawAPI) *Circle {
	return &Circle{
		Shape:  NewShape(drawAPI),
		x:      x,
		y:      y,
		radius: radius,
	}
}

// Draw 绘制圆形
func (c *Circle) Draw() {
	c.drawAPI.DrawCircle(c.x, c.y, c.radius)
}

// Rectangle 矩形
type Rectangle struct {
	*Shape
	x, y, width, height int
}

// NewRectangle 创建矩形
func NewRectangle(x, y, width, height int, drawAPI DrawAPI) *Rectangle {
	return &Rectangle{
		Shape:  NewShape(drawAPI),
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

// Draw 绘制矩形
func (r *Rectangle) Draw() {
	r.drawAPI.DrawRectangle(r.x, r.y, r.width, r.height)
}
