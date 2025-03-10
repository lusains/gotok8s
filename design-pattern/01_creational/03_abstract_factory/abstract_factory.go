package abstract_factory

// Button 按钮接口
type Button interface {
	Render() string
	OnClick()
}

// Input 输入框接口
type Input interface {
	Render() string
	OnInput(value string)
}

// GUIFactory GUI工厂接口
type GUIFactory interface {
	CreateButton() Button
	CreateInput() Input
}

// DarkButton 暗色主题按钮
type DarkButton struct{}

func (d *DarkButton) Render() string {
	return "渲染暗色按钮"
}

func (d *DarkButton) OnClick() {
	println("暗色按钮被点击")
}

// DarkInput 暗色主题输入框
type DarkInput struct{}

func (d *DarkInput) Render() string {
	return "渲染暗色输入框"
}

func (d *DarkInput) OnInput(value string) {
	println("暗色输入框输入:", value)
}

// LightButton 亮色主题按钮
type LightButton struct{}

func (l *LightButton) Render() string {
	return "渲染亮色按钮"
}

func (l *LightButton) OnClick() {
	println("亮色按钮被点击")
}

// LightInput 亮色主题输入框
type LightInput struct{}

func (l *LightInput) Render() string {
	return "渲染亮色输入框"
}

func (l *LightInput) OnInput(value string) {
	println("亮色输入框输入:", value)
}

// DarkFactory 暗色主题工厂
type DarkFactory struct{}

func (d *DarkFactory) CreateButton() Button {
	return &DarkButton{}
}

func (d *DarkFactory) CreateInput() Input {
	return &DarkInput{}
}

// LightFactory 亮色主题工厂
type LightFactory struct{}

func (l *LightFactory) CreateButton() Button {
	return &LightButton{}
}

func (l *LightFactory) CreateInput() Input {
	return &LightInput{}
}
