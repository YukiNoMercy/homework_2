package main

import (
	"fmt"
	"strings"
)

func input_string() string {
	var str string
	fmt.Println("Введите строку: ")
	fmt.Scanln(&str)
	return str
}

func rewrite_string(str string) string {
	return strings.Replace(str, "1", "one", -1)
}
func main() {

	s := input_string()
	new_s := rewrite_string(s)
	fmt.Println(new_s)

}
