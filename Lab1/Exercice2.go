package main

import (
	"fmt"
	"strings"
)

func main() {
	lineWidth := 5
	symb := "x"
	for i := 1; i <= lineWidth; i++ {
		lineSymb := strings.Repeat(symb, i)
		formatStr := fmt.Sprintf("%%%ds\n", lineWidth)
		fmt.Printf(formatStr, lineSymb)
	}
	for j := 4; j > 0; j-- {
		lineSymb := strings.Repeat(symb, j)
		formatStr := fmt.Sprintf("%%%ds\n", lineWidth)
		fmt.Printf(formatStr, lineSymb)
	}
}
