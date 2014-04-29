package main

import (
	"fmt"
	"math"
)

func main() {
	var n, r float64
	fmt.Scan(&n,&r)
	fmt.Println(2*n*r*math.Sin(math.Pi/n))
}