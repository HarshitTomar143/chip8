package chip8

type CPU struct{
	V [16]byte  // normal registers
	I uint16 // Index Register
	PC uint16 // Program Counter
	SP byte // stack Pointer
	Stack[16]uint16 //Call Stack 
}