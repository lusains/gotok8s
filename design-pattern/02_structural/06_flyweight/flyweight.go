package flyweight

// TreeType 树的类型（享元对象）
type TreeType struct {
	name    string
	color   string
	texture string
}

// NewTreeType 创建树类型
func NewTreeType(name, color, texture string) *TreeType {
	return &TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}

// Draw 绘制树
func (t *TreeType) Draw(x, y int) {
	println("绘制树 [类型:", t.name, ", 颜色:", t.color, ", 纹理:", t.texture,
		", 位置: x=", x, ", y=", y, "]")
}

// Tree 树实例（包含外部状态）
type Tree struct {
	x, y     int
	treeType *TreeType
}

// NewTree 创建树实例
func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{
		x:        x,
		y:        y,
		treeType: treeType,
	}
}

// Draw 绘制树实例
func (t *Tree) Draw() {
	t.treeType.Draw(t.x, t.y)
}

// TreeFactory 树工厂（享元工厂）
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

// NewTreeFactory 创建树工厂
func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]*TreeType),
	}
}

// GetTreeType 获取树类型（如果存在则复用，不存在则创建）
func (f *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + color + texture
	if treeType, ok := f.treeTypes[key]; ok {
		return treeType
	}
	treeType := NewTreeType(name, color, texture)
	f.treeTypes[key] = treeType
	return treeType
}

// Forest 森林（客户端）
type Forest struct {
	trees   []*Tree
	factory *TreeFactory
}

// NewForest 创建森林
func NewForest() *Forest {
	return &Forest{
		trees:   make([]*Tree, 0),
		factory: NewTreeFactory(),
	}
}

// PlantTree 种植树
func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.factory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

// Draw 绘制森林
func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}
