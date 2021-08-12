package main

import (
	"errors"
	"fmt"
)

// PC representa um computador pessoal
type PC struct {
	Motherboard string
	CPU         string
	GPU         string
	PowerSupply string
}

// pcMod é uma função que modifica o pc, ou seja, uma etapa da contrução
type pcMod func(*PC) error

// PCBuilder é o contrutor para o tipo PC
type PCBuilder struct {
	steps []pcMod //passos da contrução
}

func NewPCBuilder() *PCBuilder {
	return &PCBuilder{}
}

// Build controi o pc, aplicando cada passo da contrução ao PC
func (b *PCBuilder) Build() (*PC, error) {
	pc := &PC{}
	for _, step := range b.steps {
		if err := step(pc); err != nil {
			return nil, err
		}
	}
	return pc, nil
}

func (b *PCBuilder) WithMotherboard(motherboard string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) error {
		if motherboard == "" {
			return errors.New("Nome da placa mãe inválida.")
		}
		pc.Motherboard = motherboard
		return nil
	})
	return b
}

func (b *PCBuilder) WithCPU(cpu string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) error {
		if cpu == "" {
			return errors.New("Nome da cpu inválida.")
		}
		pc.CPU = cpu
		return nil
	})
	return b
}

func (b *PCBuilder) WithGPU(gpu string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) error {
		if gpu == "" {
			return errors.New("Nome da gpu inválida.")
		}
		pc.GPU = gpu
		return nil
	})
	return b
}

func (b *PCBuilder) WithPowerSupply(powerSupply string) *PCBuilder {
	b.steps = append(b.steps, func(pc *PC) error {
		if powerSupply == "" {
			return errors.New("Nome da fonte inválida.")
		}
		pc.PowerSupply = powerSupply
		return nil
	})
	return b
}

func main() {
	pb := NewPCBuilder()
	pc, err := pb.
		WithMotherboard("Gigabyte").
		WithCPU("Ryzen 7 3700").
		WithGPU("NVidia RTX 2060").
		WithPowerSupply("G750H").
		Build()

	if err != nil {
		fmt.Println("erro ao contruir pc: ", err)
	} else {
		fmt.Println(pc)
	}

	pc, err = pb.
		WithMotherboard("Gigabyte").
		WithCPU("").
		WithGPU("").
		WithPowerSupply("G750H").
		Build()

	if err != nil {
		fmt.Println("erro ao contruir pc: ", err)
	} else {
		fmt.Println(pc)
	}

}
