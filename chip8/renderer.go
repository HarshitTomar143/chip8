package chip8

import(
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth= 64
	screenHeight= 32
	scale= 10
)

type Game struct{
	Emu *Emulator
}

func (g *Game) Update() error{
	// Run multiple CPU cycles per frame
	for i := 0; i<10; i++ {
		g.Emu.Cycle()
		g.Emu.UpdateTimers()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	for y := 0; y < 32; y++ {
		for x := 0; x < 64; x++ {
			if g.Emu.Display[y][x] == 1 {
				for dy := 0; dy < scale; dy++ {
					for dx := 0; dx < scale; dx++ {
						screen.Set(
							x*scale+dx,
							y*scale+dy,
							color.White,
						)
					}
				}
			}
		}
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth * scale, screenHeight * scale
}