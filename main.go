package main

import (
	"chip8/chip8"
	"fmt"
	"log"
	// "time"
)

func main() {
	emu := chip8.New()

	err := emu.LoadRom("roms/test.ch8")

	if err != nil {
		log.Fatal(err)
	}

	// Run a few cycles for now

	for i := 0; i < 10; i++ {
		emu.Cycle()
		emu.UpdateTimers()
		// time.Sleep(time.Millisecond * 5) // we have implemented this to check whether delaytimer is matchin up with cpu or not as cpu is very fast so we have to slow it down to run the test files
		fmt.Printf("PC: %04X V0: %d\n", emu.CPU.PC, emu.CPU.V[0])
	}

	fmt.Printf("Font[0]: %02X %02X %02X %02X %02X\n",
		emu.Memory.Data[0x050],
		emu.Memory.Data[0x051],
		emu.Memory.Data[0x052],
		emu.Memory.Data[0x053],
		emu.Memory.Data[0x054],
	)

}
