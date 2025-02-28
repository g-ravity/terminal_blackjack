package main

import (
	"fmt"
	"time"
)

func printWithTypingEffect(text string, delay int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
