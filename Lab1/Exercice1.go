package main

import (
	"fmt"
	"math"
)

func floorAndCeil(in float64) (int, int) {
	return int(math.Floor(in)), int(math.Ceil(in))
}
func main() {
	fmt.Println(floorAndCeil(10.4))
}
