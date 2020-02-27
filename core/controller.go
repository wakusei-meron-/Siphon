package core

import "github.com/eiannone/keyboard"

func MustSetup() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
}
