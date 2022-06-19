package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*通过下面的我们可以发现 len是字节长度*/
	// 通过将字符串封装成[]rune() 切片或者使用utf8.RunneCOuntInString 才能获取字符串的真正的长度
	str := "你好 china"

	//golang中string底层是通过byte数组实现的，座椅直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))

	//以下两种都可以得到str的字符串长度
	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	//通过rune类型处理unicode字符
	fmt.Println("rune:", len([]rune(str)))
	//
	//len(str): 12
	//RuneCountInString: 8
	//rune: 8

}
