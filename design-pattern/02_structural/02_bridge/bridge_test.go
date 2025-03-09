package bridge

import "testing"

func TestBridge(t *testing.T) {
	redCircle := NewCircle(100, 100, 10, &RedCircle{})
	greenCircle := NewCircle(100, 100, 10, &GreenCircle{})
	redRectangle := NewRectangle(100, 100, 50, 30, &RedCircle{})
	greenRectangle := NewRectangle(100, 100, 50, 30, &GreenCircle{})

	// 测试不同颜色的圆形
	redCircle.Draw()
	greenCircle.Draw()

	// 测试不同颜色的矩形
	redRectangle.Draw()
	greenRectangle.Draw()

	// 验证桥接模式的结构
	if _, ok := redCircle.drawAPI.(*RedCircle); !ok {
		t.Error("redCircle应该使用RedCircle实现")
	}

	if _, ok := greenRectangle.drawAPI.(*GreenCircle); !ok {
		t.Error("greenRectangle应该使用GreenCircle实现")
	}
}
