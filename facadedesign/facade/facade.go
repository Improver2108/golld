package facade

import (
	"fmt"

	"github.com/improver2108/golld/facadedesign/facade/internal/computer"
)

type computerFacade struct {
	powerSupply     *computer.PowerSupply
	coolingSystem   *computer.CoolingSystem
	cpu             *computer.CPU
	memory          *computer.Memory
	hardDrive       *computer.HardDrive
	bios            *computer.Bios
	operatingSystem *computer.OperatingSystem
}

func NewComputerFacade() *computerFacade {
	return &computerFacade{
		powerSupply:     computer.NewPowerSupply(),
		coolingSystem:   computer.NewCoolingSystem(),
		cpu:             computer.NewCPU(),
		memory:          computer.NewMemory(),
		hardDrive:       computer.NewHardDrive(),
		bios:            computer.NewBIOS(),
		operatingSystem: computer.NewOperatingSystem(),
	}
}

func (f *computerFacade) StartComputer() {
	fmt.Println("-------Starting Computer-------")
	f.powerSupply.ProvidePower()
	f.coolingSystem.StartFans()
	f.bios.Boot(f.cpu, f.memory)
	f.hardDrive.SpinUp()
	f.operatingSystem.Load()
	fmt.Println("Computer booted successfully")
}
