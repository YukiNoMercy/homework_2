package main

import (
	"fmt"
	"math"
)

func hipot() {
	var a, b int
	fmt.Println("Введите длины катетов гиппотенузы: ")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	fmt.Println(math.Hypot(float64(a), float64(b)))
}
func main() {
	hipot()
}
