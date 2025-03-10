package facade

// 子系统1：CPU
type CPU struct{}

func (c *CPU) Freeze() {
	println("CPU: 冻结...")
}

func (c *CPU) Jump(position string) {
	println("CPU: 跳转到", position)
}

func (c *CPU) Execute() {
	println("CPU: 执行指令...")
}

// 子系统2：内存
type Memory struct{}

func (m *Memory) Load(position string, data string) {
	println("内存: 从", position, "加载数据", data)
}

// 子系统3：硬盘
type HardDrive struct{}

func (h *HardDrive) Read(lba int, size int) string {
	println("硬盘: 读取扇区", lba, "大小", size)
	return "数据"
}

// 外观
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Start 启动计算机
func (c *ComputerFacade) Start() {
	c.cpu.Freeze()
	c.memory.Load("0x00", c.hardDrive.Read(0, 1024))
	c.cpu.Jump("0x00")
	c.cpu.Execute()
}

// 子系统4：显示器
type Display struct{}

func (d *Display) TurnOn() {
	println("显示器: 开启...")
}

func (d *Display) TurnOff() {
	println("显示器: 关闭...")
}

// 子系统5：操作系统
type OS struct{}

func (o *OS) LoadOS() {
	println("操作系统: 加载中...")
}

// 扩展的外观
type ExtendedComputerFacade struct {
	*ComputerFacade
	display *Display
	os      *OS
}

func NewExtendedComputerFacade() *ExtendedComputerFacade {
	return &ExtendedComputerFacade{
		ComputerFacade: NewComputerFacade(),
		display:        &Display{},
		os:             &OS{},
	}
}

// StartComputer 完整的启动计算机流程
func (e *ExtendedComputerFacade) StartComputer() {
	e.display.TurnOn()
	e.Start()
	e.os.LoadOS()
	println("计算机启动完成！")
}

// ShutDown 关闭计算机
func (e *ExtendedComputerFacade) ShutDown() {
	println("开始关闭计算机...")
	e.display.TurnOff()
	println("计算机已关闭！")
}
