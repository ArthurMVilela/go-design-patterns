package main

import "fmt"

type PC struct {
	Motherboard string
	CPU         string
	GPU         string
	PowerSupply string
}

type PCBuilder struct {
	pc *PC
}

func NewPCBuilder() *PCBuilder {
	return &PCBuilder{&PC{}}
}

func (b *PCBuilder) WithMotherboard(motherboard string) *PCBuilder {
	b.pc.Motherboard = motherboard
	return b
}

func (b *PCBuilder) WithCPU(cpu string) *PCBuilder {
	b.pc.CPU = cpu
	return b
}

func (b *PCBuilder) WithGPU(gpu string) *PCBuilder {
	b.pc.GPU = gpu
	return b
}

func (b *PCBuilder) WithPowerSupply(powerSupply string) *PCBuilder {
	b.pc.PowerSupply = powerSupply
	return b
}

func (b *PCBuilder) Build() *PC {
	return b.pc
}

func main() {
	pb := NewPCBuilder()
	pc := pb.
		WithMotherboard("Gigabyte").
		WithCPU("Ryzen 7 3700").
		WithGPU("NVidia RTX 2060").
		WithPowerSupply("G750H").
		Build()

	fmt.Println(pc)
}
