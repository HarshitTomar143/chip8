package chip8

import (
	"fmt"
)



func (e *Emulator)Cycle (){

	// DEBUG: print state BEFORE execution
	fmt.Printf(
		"PC=%04X OPCODE=%04X V0=%d V1=%d I=%04X\n",
		e.CPU.PC,
		uint16(e.Memory.Data[e.CPU.PC])<<8|uint16(e.Memory.Data[e.CPU.PC+1]),
		e.CPU.V[0],
		e.CPU.V[1],
		e.CPU.I,
	)

	//Fetch
	opcode := uint16(e.Memory.Data[e.CPU.PC])<<8 | uint16(e.Memory.Data[e.CPU.PC+1])
	e.CPU.PC += 2

	//Decode
	X := (opcode & 0x0F00) >> 8
	nn:= byte(opcode & 0x00FF)
	nnn := opcode & 0x0FFF
	Y := (opcode & 0X00F0) >> 8

	// Executing (start small)
	switch opcode & 0xF000 {

	case 0x1000:
		e.CPU.PC= nnn

	case 0x3000:
		if e.CPU.V[X] == nn{
			e.CPU.PC += 2
		}

	case 0x4000: 
		if e.CPU.V[X] != nn {
			e.CPU.PC += 2
		}

	case 0x5000:
		if opcode&0x000F == 0 {
			if e.CPU.V[X] == e.CPU.V[Y] {
				e.CPU.PC += 2
			}
		}
	
	case 0x9000:
		if opcode&0x000F == 0{
			if e.CPU.V[X] != e.CPU.V[Y]{
				e.CPU.PC += 2
			}
		}

	case 0x6000:
		e.CPU.V[X]= nn
		
	case 0x7000:
		e.CPU.V[X] += nn
		
	case 0xA000:
		e.CPU.I= nnn
		
	default:
		fmt.Printf("Unknown OPcode : %04X\n",opcode)	
	}
}