package chip8
import (
	"time"
	// "fmt"
)

type Emulator struct{
	CPU CPU
	Memory Memory
	Debug bool

	lastTimerTick time.Time // timer

	Display [32][64]byte // [Y][X], 0 or 1
}

func New() *Emulator{
	e := &Emulator{}
	e.CPU.PC = 0x200 // program got started here 
	e.lastTimerTick= time.Now()
	e.ClearDisplay()

	// Load fontset into memory starting at 0x050
	for i := 0; i< len(chip8Font); i++ {
		e.Memory.Data[0x050+i] = chip8Font[i]
	} 

	return e
}

func (e *Emulator) UpdateTimers(){
	now := time.Now()
	if now.Sub(e.lastTimerTick) >= time.Second/60{
		if e.CPU.DelayTimer > 0 {
			e.CPU.DelayTimer--
		}
		if e.CPU.SoundTimer > 0 {
			e.CPU.SoundTimer--
		}
		e.lastTimerTick = now
	}
}

func (e *Emulator) ClearDisplay(){
	for y := 0; y<32; y++ {
		for x := 0; x<64; x++ {
			e.Display[y][x]= 0
			
		}
	}
}

