package composite

// Component 组件接口
type Component interface {
	Add(Component)
	Remove(Component)
	GetChild(int) Component
	Operation() string
	GetName() string
}

// Composite 组合类（如文件夹）
type Composite struct {
	name     string
	children []Component
}

func NewComposite(name string) *Composite {
	return &Composite{
		name:     name,
		children: make([]Component, 0),
	}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	for i, child := range c.children {
		if child.GetName() == component.GetName() {
			c.children = append(c.children[:i], c.children[i+1:]...)
			break
		}
	}
}

func (c *Composite) GetChild(index int) Component {
	if index >= 0 && index < len(c.children) {
		return c.children[index]
	}
	return nil
}

func (c *Composite) Operation() string {
	result := "Folder(" + c.name + ") ["
	for i, child := range c.children {
		if i > 0 {
			result += " "
		}
		result += child.Operation()
	}
	result += "]"
	return result
}

func (c *Composite) GetName() string {
	return c.name
}

// Leaf 叶子类（如文件）
type Leaf struct {
	name string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

func (l *Leaf) Add(Component) {
	// 叶子节点不能添加子节点
}

func (l *Leaf) Remove(Component) {
	// 叶子节点不能删除子节点
}

func (l *Leaf) GetChild(index int) Component {
	// 叶子节点没有子节点
	return nil
}

func (l *Leaf) Operation() string {
	return "File(" + l.name + ")"
}

func (l *Leaf) GetName() string {
	return l.name
}
