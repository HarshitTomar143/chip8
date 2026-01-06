package main

import (
	"chip8/chip8"
	// "fmt"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	// "time"
)

func main() {
	emu := chip8.New()

	err := emu.LoadRom("roms/test.ch8")

	if err != nil {
		log.Fatal(err)
	}

	game := &chip8.Game{
		Emu : emu,
	}

	ebiten.SetWindowSize(64*10, 32*10)
	ebiten.SetWindowTitle("CHIP-8 Emulator")

	if err := ebiten.RunGame(game);	err != nil{
		log.Fatal(err)
	}

	// Run a few cycles for now

	// for i := 0; i < 10; i++ {
	// 	emu.Cycle()
	// 	emu.UpdateTimers()
	// 	// time.Sleep(time.Millisecond * 5) // we have implemented this to check whether delaytimer is matchin up with cpu or not as cpu is very fast so we have to slow it down to run the test files
	// 	fmt.Printf("PC: %04X V0: %d\n", emu.CPU.PC, emu.CPU.V[0])
	// }

	// fmt.Println("Framebuffer output:")
	// for y := 5; y < 10; y++ {
	// 	for x := 10; x < 18; x++ {
	// 		fmt.Print(emu.Display[y][x])
	// 	}
	// 	fmt.Println()
	// }


}
