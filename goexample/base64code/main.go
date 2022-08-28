package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "advadsf333333333333333333333ffffffffffffffffffffffffffffffff3"
	toString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(string(toString))
	decodeString, _ := base64.StdEncoding.DecodeString((string(toString)))
	fmt.Println(string(decodeString))
}
