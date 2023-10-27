package main

import "fmt"

func input_matric(size int) [][]int {
	var input int
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	fmt.Println("Введите матрицу")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Scan(&input)
			matrix[i][j] = input
		}
	}
	return matrix
}

func output_matrix(matrix [][]int, n int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
func main() {
	var matrix_size int
	fmt.Println("Введите размер матрицы")
	fmt.Scan(&matrix_size)
	matrix := input_matric(matrix_size)

	output_matrix(matrix, matrix_size)
}
