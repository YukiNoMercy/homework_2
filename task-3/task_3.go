package main

import "fmt"

func input_array(array *[]int, size int) {
	fmt.Println("Введите массив")
	temp := 0
	for i := 0; i < size; i++ {
		fmt.Scan(&temp)
		*array = append(*array, temp)
	}
}

func cycle_mov(array *[]int, mov int) {
	last_el := 0
	temp_ar := *array
	for i := 0; i < mov; i++ {
		last_el = temp_ar[len(temp_ar)-1]
		for i := len(temp_ar) - 1; i > 0; i-- {
			temp_ar[i] = temp_ar[i-1]
		}
		temp_ar[0] = last_el
		fmt.Println(temp_ar, " - После сдвига", i+1)
	}
}
func main() {
	var N, move int
	fmt.Println("Введите размер массива: ")
	fmt.Scan(&N)
	var array []int
	input_array(&array, N)
	fmt.Println("Введите число сдвигов: ")
	fmt.Scan(&move)
	cycle_mov(&array, move)
}
