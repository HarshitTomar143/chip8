package main

import "os"

func main() {
	data := []byte{
		0x60, 0x0A,
		0x70, 0x05,
		0x12, 0x00,
	}
	os.WriteFile("roms/test.ch8", data, 0644)
}
