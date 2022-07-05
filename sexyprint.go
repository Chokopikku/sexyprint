package sexyprint

import (
	"fmt"
	"time"
)

func Clear() {
	// Clear terminal
	fmt.Print("\033[2J")
	// Move to position 0:0
	fmt.Print("\033[H")
}

func DelayedPrint(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

func Ellipsis() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Print(".")
		time.Sleep(time.Millisecond * 500)
		fmt.Print(".")
		time.Sleep(time.Millisecond * 500)
		fmt.Print(".")
		time.Sleep(time.Millisecond * 500)
		fmt.Print("\b\b\b   \b\b\b")
	}
	fmt.Println()
}
