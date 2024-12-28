package utils

import (
	"fmt"
	"time"
)

// ShowProgressBar displays a simple progress bar
func ShowProgressBar(total int) {
	for i := 0; i <= total; i++ {
		fmt.Printf("\rProgress: [%s%s] %d%%", string(replicate('#', i)), string(replicate(' ', total-i)), i*100/total)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println()
}

// replicate returns a string of repeated characters
func replicate(char rune, count int) string {
	return string([]rune{char}[0:count])
}
