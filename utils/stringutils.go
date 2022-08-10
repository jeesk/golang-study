package utils

import "fmt"

const name = "asdfadsf"

func init() {
	fmt.Println("uitl init ", name, name1)
}

var name1 = "name1"

func GetLength(str string) int {

	return len(str)
}
