package input

import (
	"github.com/eiannone/keyboard"
)

func InitKeyboard() error {
	return keyboard.Open()
}

func CloseKeyboard()  {
	keyboard.Close()
}

func ReadKey() (rune, bool) {
	if char, key, err := keyboard.GetKey(); err == nil {
		if key == keyboard.KeyEsc {
			return 'q', true
		}
		return char, false
	}
	return 0, false
}