package builder

// Computer 电脑产品
type Computer struct {
	CPU    string
	Memory string
	Disk   string
	GPU    string
}

// ComputerBuilder 电脑建造者接口
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetMemory(memory string) ComputerBuilder
	SetDisk(disk string) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	Build() *Computer
}

// GamingComputerBuilder 游戏电脑建造者
type GamingComputerBuilder struct {
	computer *Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{
		computer: &Computer{},
	}
}

func (b *GamingComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *GamingComputerBuilder) SetMemory(memory string) ComputerBuilder {
	b.computer.Memory = memory
	return b
}

func (b *GamingComputerBuilder) SetDisk(disk string) ComputerBuilder {
	b.computer.Disk = disk
	return b
}

func (b *GamingComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *GamingComputerBuilder) Build() *Computer {
	return b.computer
}

// Director 指挥者
type Director struct {
	builder ComputerBuilder
}

func NewDirector(builder ComputerBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// ConstructGamingComputer 构建游戏电脑
func (d *Director) ConstructGamingComputer() *Computer {
	return d.builder.
		SetCPU("Intel i9").
		SetMemory("32GB").
		SetDisk("2TB SSD").
		SetGPU("RTX 3080").
		Build()
}

// ConstructOfficeComputer 构建办公电脑
func (d *Director) ConstructOfficeComputer() *Computer {
	return d.builder.
		SetCPU("Intel i5").
		SetMemory("16GB").
		SetDisk("512GB SSD").
		SetGPU("Integrated Graphics").
		Build()
}
