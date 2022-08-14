package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str = "你好hello"
	fmt.Println(len(str))
	// 使用rune 转换成int32
	print(len([]rune(str)))
	fmt.Println()
	println(utf8.RuneCountInString(str))

	for i := 0; i < len(str); i++ {
		fmt.Println("%x", str[i])
	}
	fmt.Println()
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at %d  ", runeValue, i)
		w = width
	}

}
