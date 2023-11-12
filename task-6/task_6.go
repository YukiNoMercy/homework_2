package main

import (
	"bufio"
	"fmt"
	"os"
)

func input_date(input *[]string) {
	fmt.Println("Введите числа:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num := scanner.Text()
	for i := 0; i < len(num); i++ {
		*input = append(*input, string(num[i]))
	}
}

func add_dict(dict *map[string]bool, data []string) {
	temp_dict_link := make(map[string]bool)
	for i := 0; i < len(data); i++ {
		_, found := temp_dict_link[string(data[i])]
		if !found {
			temp_dict_link[string(data[i])] = true
		}
	}
	*dict = temp_dict_link

}

func main() {

	var data []string
	input_date(&data)

	var slov map[string]bool
	add_dict(&slov, data)

	fmt.Println(len(slov) - 1)

}
