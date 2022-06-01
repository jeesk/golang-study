package main

import (
	"fmt"
)

func main() {
	friends, err := Find("附近的人",
		WithSex(1),
		WithAge(30),
		WithHeight(160),
		WithWeight(55),
		WithHobby("爬山"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(friends)
}
