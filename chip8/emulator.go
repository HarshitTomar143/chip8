package chip8

type Emulator struct{
	CPU CPU
	Memory Memory
	Debug bool
}

func New() *Emulator{
	e := &Emulator{}
	e.CPU.PC = 0x200 // program got started here 
	return e
}