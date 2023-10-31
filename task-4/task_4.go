package main

import "fmt"

type matric [][]int

func input_matric(size int) matric {
	var input int
	matrix := make(matric, size)
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

func output_matrix(matrix matric) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func is_simetri(matrix matric) string {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[i][j] != matrix[j][i] {
				return "no"
			}
		}
	}
	return "yes"
}
func main() {
	var matrix_size int
	fmt.Println("Введите размер матрицы")
	fmt.Scan(&matrix_size)
	matrix := input_matric(matrix_size)

	// output_matrix(matrix)

	simetry := is_simetri(matrix)
	fmt.Println(simetry)
}
