package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"time"
	"unicode"
)

func rgb(text []rune, i int) (int, int, int) {
	var f = 0.005
	bounds := func(rgbVal float64) int {
		return int(math.Max(130, (rgbVal*127 + 128)))
	}
	return bounds(math.Sin(f*float64(i) + 0)),
		bounds(math.Sin(f*float64(i) + 2*math.Pi/3)),
		bounds(math.Sin(f*float64(i) + 4*math.Pi/3))
}

func print(output []rune) {
	for i := 0; i < len(output); i++ {
		char := output[i]
		r, g, b := rgb(output, i)
		if unicode.IsSpace(char) {
			time.Sleep(80 * time.Millisecond)
		}
		if unicode.IsNumber(char) || unicode.IsPunct(char) {
			r, g, b = 255, 255, 255
		}
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m",
			r, g, b, char)
	}
	fmt.Println()
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gogocat")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	print(output)
}
