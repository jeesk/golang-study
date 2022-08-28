package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		ucl := strings.ToUpper(scaner.Text())
		fmt.Println(ucl)
	}
}
