package main
import (
	"fmt"
	"log"
	"chip8/chip8"
)

func main(){
	emu := chip8.New()

	err := emu.LoadRom("roms/test.ch8")
	
	if err != nil{
		log.Fatal(err) 
	}

	// Run a few cycles for now

	for i := 0; i<10 ; i++ {
		emu.Cycle()
		fmt.Printf("PC: %04X V0: %d\n", emu.CPU.PC, emu.CPU.V[0])
	}
}

