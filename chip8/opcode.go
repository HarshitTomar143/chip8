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
	Y := (opcode & 0X00F0) >> 4

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

	case 0x2000: //call Address
		e.CPU.Stack[e.CPU.SP]= e.CPU.PC
		e.CPU.SP++
		e.CPU.PC= nnn
	
	case 0x0000: //RET
		switch opcode {
		case 0x00EE:
			e.CPU.SP--
			e.CPU.PC = e.CPU.Stack[e.CPU.SP]
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

	//Alu codes (this is the largest fmaily before timers, graphics etc)
	case 0x8000:
		switch opcode & 0x000F {
		case 0x0: // LD Vx, Vy
			e.CPU.V[X] = e.CPU.V[Y]

		case 0x1: //or
			e.CPU.V[X] |= e.CPU.V[Y]
			
		case 0x2: 	// and
			e.CPU.V[X] &= e.CPU.V[Y]

		case 0x3: // XOR
			e.CPU.V[X] ^= e.CPU.V[Y]
			
		case 0x4: // ADD with carry
			sum := uint16(e.CPU.V[X]) + uint16(e.CPU.V[Y])
			if sum> 0xFF {
				e.CPU.V[0xF] = 1
			}else{
				e.CPU.V[0xF] = 0
			}
			e.CPU.V[X] = byte(sum)	
			

		case 0x5: //Sub Vx= Vx - Vy
			if e.CPU.V[X] >= e.CPU.V[Y]{
				e.CPU.V[0XF] = 1
			}else{
				e.CPU.V[0xF]= 0
			}
			e.CPU.V[X] -= e.CPU.V[Y]

		case 0x6: // SHR
			e.CPU.V[0xF] = e.CPU.V[X] & 0x1
			e.CPU.V[X] >>= 1

		case 0x7: // SUBN Vx = Vy - Vx
			if e.CPU.V[Y] >= e.CPU.V[X]{
				e.CPU.V[0xF] = 1
			}else{
				e.CPU.V[0xF] = 0
			}	
			e.CPU.V[X] = e.CPU.V[Y] - e.CPU.V[X]
		
		case 0xE: //SHL
			e.CPU.V[0xF] = (e.CPU.V[X] >> 7) & 0x1
			e.CPU.V[X] <<= 1
		}
				
		
	default:
		fmt.Printf("Unknown OPcode : %04X\n",opcode)	
	}
}