package main

import "untitled/core"

func main() {
    p := core.NewPeaberry()
    p.Run()
	//tm.Clear() // Clear current screen
	//
	//
	//err := keyboard.Open()
	//if err != nil {
	//	panic(err)
	//}
	//defer keyboard.Close()
	//
	//x, y := 1, 3
	//
	//for {
	//	tm.Clear()
	//	// By moving cursor to top-left position we ensure that console output
	//	// will be overwritten each time, instead of adding new.
	//	tm.MoveCursor(1, 1)
	//
	//	tm.Println("Current Time:", time.Now().Format(time.RFC1123))
	//
	//	tm.Flush() // Call it every time at the end of rendering
	//
	//	char, key, err := keyboard.GetKey()
	//	if (err != nil) {
	//		panic(err)
	//	} else if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC || char == 'q' {
	//		return
	//	}
	//	tm.MoveCursor(1, 2)
	//	tm.Printf("You pressed: %q, x:%d, y:%d", char, x, y)
	//
	//	switch char {
	//	case 'h':
	//		if x > 1 {x -= 1}
	//	case 'j':
	//		if y < tm.Height() {y += 1}
	//	case 'k':
	//		if y > 1 {y -= 1}
	//	case 'l':
	//		if x < tm.Width() {x += 1}
	//	}
	//	tm.MoveCursor(x, y)
	//	tm.Print("*")
	//}
}
