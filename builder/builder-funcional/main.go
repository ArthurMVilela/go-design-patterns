package main

import "fmt"

// PC representa um computador pessoal
type PC struct {
	Motherboard string
	CPU         string
	GPU         string
	PowerSupply string
}

// pcMod é uma função que modifica o pc, ou seja, uma etapa da contrução
type pcMod func(*PC)

// PCBuilder é o contrutor para o tipo PC
type PCBuilder struct {
	steps []pcMod //passos da contrução
}

func NewPCBuilder() *PCBuilder {
	return &PCBuilder{}
}

// Build controi o pc, aplicando cada passo da contrução ao PC
func (b *PCBuilder) Build() *PC {
	pc := &PC{}
	for _, step := range b.steps {
		step(pc)
	}
	return pc
}

func (b *PCBuilder) WithMotherboard(motherboard string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) {
		pc.Motherboard = motherboard
	})
	return b
}

func (b *PCBuilder) WithCPU(cpu string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) {
		pc.CPU = cpu
	})
	return b
}

func (b *PCBuilder) WithGPU(gpu string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) {
		pc.GPU = gpu
	})
	return b
}

func (b *PCBuilder) WithPowerSupply(powerSupply string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) {
		pc.PowerSupply = powerSupply
	})
	return b
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
