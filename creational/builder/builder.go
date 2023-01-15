package builder

// struct to build

type PersonalComputer struct {
	ramCap int
	hddCap int
	cpu    string
	os     string
	gpu    string
}

// builder need a interface

type PCBuilder interface {
	SetRAM(ram int) PCBuilder
	SetCPU(cpu string) PCBuilder
	SetHDD(hdd int) PCBuilder
	SetOS(os string) PCBuilder
	SetGPU(gpu string) PCBuilder
}

type HomeEditionBuilder struct {
	pc PersonalComputer
}

func (builder *HomeEditionBuilder) SetRAM(ram int) PCBuilder {
	builder.pc.ramCap = ram
	return builder
}

func (builder *HomeEditionBuilder) SetHDD(hdd int) PCBuilder {
	builder.pc.hddCap = hdd
	return builder
}

func (builder *HomeEditionBuilder) SetCPU(cpu string) PCBuilder {
	builder.pc.cpu = cpu
	return builder
}

func (builder *HomeEditionBuilder) SetOS(os string) PCBuilder {
	builder.pc.os = os
	return builder
}

func (builder *HomeEditionBuilder) SetGPU(gpu string) PCBuilder {
	builder.pc.gpu = gpu
	return builder
}

func (builder *HomeEditionBuilder) GetPC() PersonalComputer {
	return builder.pc
}

type GameEditionPCBuilder struct {
	pc PersonalComputer
}

func (builder *GameEditionPCBuilder) SetRAM(ram int) PCBuilder {
	builder.pc.ramCap = ram
	return builder
}

func (builder *GameEditionPCBuilder) SetHDD(hdd int) PCBuilder {
	builder.pc.hddCap = hdd
	return builder
}

func (builder *GameEditionPCBuilder) SetCPU(cpu string) PCBuilder {
	builder.pc.cpu = cpu
	return builder
}

func (builder *GameEditionPCBuilder) SetOS(os string) PCBuilder {
	builder.pc.os = os
	return builder
}

func (builder *GameEditionPCBuilder) SetGPU(gpu string) PCBuilder {
	builder.pc.gpu = gpu
	return builder
}

func (builder *GameEditionPCBuilder) GetPC() PersonalComputer {
	return builder.pc
}

// manufacture will define builder needed

type Manufacture struct {
	builder PCBuilder
}

func (m *Manufacture) SetBuilder(builder PCBuilder) {
	m.builder = builder
}

// func (m *Manufacture) ConstructPC() {
// 	m.builder.SetCPU("i4").SetHDD(500).SetRAM().SetOS().SetGPU()
// }
