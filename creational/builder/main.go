package main

import (
	"fmt"
)

// Component: Builder interface.
// Every concrete builder should implement it.
type PCBuilder interface {
	SetProcessor()
	SetMainboard()
	SetMemorySizeGB()
	SetStorageSizeGB()
	Build() PC
}

// Component: Concrete Product
type PC struct {
	processor     string
	mainboard     string
	memorySizeGB  int32
	storageSizeGB int32
}

func (pc PC) String() string {
	spec := ""
	spec = spec + fmt.Sprintln("Processor: ", pc.processor)
	spec = spec + fmt.Sprintln("Mainboard: ", pc.mainboard)
	spec = spec + fmt.Sprintf("Memory: %v GB\n", pc.memorySizeGB)
	spec = spec + fmt.Sprintf("Storage: %v GB\n", pc.storageSizeGB)

	return spec
}

// Component: Concrete builder 1
type RyzenPCBuilder struct {
	RyzenPC PC
}

func NewRyzenPCBuilder() *RyzenPCBuilder {
	return &RyzenPCBuilder{}
}

func (r *RyzenPCBuilder) SetProcessor() {
	r.RyzenPC.processor = "AMD Ryzen 5600g"
}

func (r *RyzenPCBuilder) SetMainboard() {
	r.RyzenPC.mainboard = "MSI B550 Tomahawk"
}

func (r *RyzenPCBuilder) SetMemorySizeGB() {
	r.RyzenPC.memorySizeGB = 16
}

func (r *RyzenPCBuilder) SetStorageSizeGB() {
	r.RyzenPC.storageSizeGB = 256
}

func (r *RyzenPCBuilder) Build() PC {
	return PC{
		processor:     r.RyzenPC.processor,
		mainboard:     r.RyzenPC.mainboard,
		memorySizeGB:  r.RyzenPC.memorySizeGB,
		storageSizeGB: r.RyzenPC.storageSizeGB,
	}
}

// Component: Concrete builder 2
type IntelPCBuilder struct {
	IntelPC PC
}

func NewIntelPCBuilder() *IntelPCBuilder {
	return &IntelPCBuilder{}
}

func (r *IntelPCBuilder) SetProcessor() {
	r.IntelPC.processor = "Intel 12th Gen 12600k"
}

func (r *IntelPCBuilder) SetMainboard() {
	r.IntelPC.mainboard = "Gigabyte Z690"
}

func (r *IntelPCBuilder) SetMemorySizeGB() {
	r.IntelPC.memorySizeGB = 32
}

func (r *IntelPCBuilder) SetStorageSizeGB() {
	r.IntelPC.storageSizeGB = 512
}

func (r *IntelPCBuilder) Build() PC {
	return PC{
		processor:     r.IntelPC.processor,
		mainboard:     r.IntelPC.mainboard,
		memorySizeGB:  r.IntelPC.memorySizeGB,
		storageSizeGB: r.IntelPC.storageSizeGB,
	}
}

// Component: Director
type Director struct {
	builder PCBuilder
}

func NewDirector(b PCBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b PCBuilder) {
	d.builder = b
}

func (d *Director) BuildPC() PC {
	d.builder.SetProcessor()
	d.builder.SetMainboard()
	d.builder.SetMemorySizeGB()
	d.builder.SetStorageSizeGB()
	return d.builder.Build()
}

func main() {
	fmt.Println("\t Design Patterns in Golang(1)")
	fmt.Println("\t \t Builder Pattern")

	ryzenBuilder := NewRyzenPCBuilder()
	director := NewDirector(ryzenBuilder)
	ryzenPC := director.BuildPC()
	fmt.Println(ryzenPC)

	intelBuilder := NewIntelPCBuilder()
	director.SetBuilder(intelBuilder)
	intelPC := director.BuildPC()
	fmt.Println(intelPC)
}
