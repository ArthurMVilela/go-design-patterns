package main

import "fmt"

// PC representa um computador pessoal
type PC struct {
	Motherboard string
	CPU         string
	GPU         string
	PowerSupply string
}

// PCBuilder é o builder para o tipo PC
type PCBuilder struct {
	pc *PC
}

func NewPCBuilder() *PCBuilder {
	return &PCBuilder{&PC{}}
}

// WithMotherboard define a placa-mãe do PC
func (b *PCBuilder) WithMotherboard(motherboard string) *PCBuilder {
	b.pc.Motherboard = motherboard
	return b
}

// WithCPU define a CPU do PC
func (b *PCBuilder) WithCPU(cpu string) *PCBuilder {
	b.pc.CPU = cpu
	return b
}

// WithGPU define a GPU do PC
func (b *PCBuilder) WithGPU(gpu string) *PCBuilder {
	b.pc.GPU = gpu
	return b
}

// WithPowerSupply define a fonte do PC
func (b *PCBuilder) WithPowerSupply(powerSupply string) *PCBuilder {
	b.pc.PowerSupply = powerSupply
	return b
}

// Build controi o pc (retorna o pc construido nas estapas executadas)
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
