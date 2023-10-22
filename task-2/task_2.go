package main

import "fmt"

func is_triengle() {
	var a, b, c int
	fmt.Println("Введите стороны треугольника: ")
	fmt.Scan(&a, &b, &c)
	condition := ((a+b > c) && (b+c > a) && (a+c > b))
	if condition {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	is_triengle()

}
