package computer

import "fmt"

type PowerSupply struct{}

func NewPowerSupply() *PowerSupply {
	return &PowerSupply{}
}

func (p *PowerSupply) ProvidePower() {
	fmt.Println("Power Supply: Providing power supply")
}

type CoolingSystem struct{}

func NewCoolingSystem() *CoolingSystem {
	return &CoolingSystem{}
}

func (c *CoolingSystem) StartFans() {
	fmt.Println("Cooling System: Fans started...")
}

type CPU struct{}

func NewCPU() *CPU {
	return &CPU{}
}

func (c *CPU) initialize() {
	fmt.Println("CPU: Initialization started...")
}

type Memory struct{}

func NewMemory() *Memory {
	return &Memory{}
}

func (c *Memory) selfTest() {
	fmt.Println("Memory: Self-test passed...")
}

type HardDrive struct{}

func NewHardDrive() *HardDrive {
	return &HardDrive{}
}

func (c *HardDrive) SpinUp() {
	fmt.Println("Hard Drive: Spinning up...")
}

type Bios struct{}

func NewBIOS() *Bios {
	return &Bios{}
}

func (c *Bios) Boot(cpu *CPU, memory *Memory) {
	fmt.Println("BIOS: Booting CPU and Memory checks...")
	cpu.initialize()
	memory.selfTest()
}

type OperatingSystem struct{}

func NewOperatingSystem() *OperatingSystem {
	return &OperatingSystem{}
}

func (c *OperatingSystem) Load() {
	fmt.Println("Operating System: Loading into memory...")
}
