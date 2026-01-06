package chip8

import (
	"os"
)

func (e *Emulator) LoadRom(path string) error{
	data,err := os.ReadFile(path)
	if err != nil{
		return err
	}

	for i := 0; i< len(data); i++{
		e.Memory.Data[0x200+i]= data[i]
	}
	return nil
}