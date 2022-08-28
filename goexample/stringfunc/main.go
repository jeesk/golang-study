package main

import (
	"fmt"
	s "strings"
)

func main() {
	fmt.Println("Contains", s.Contains("hello", "h"))
	fmt.Println("Count", s.Count("test", "s"))
	fmt.Println("HasPrefix", s.HasPrefix("file.json", "file"))
	fmt.Println("Index", s.Index("test", "e"))
	fmt.Println("Join", s.Join([]string{"hello", "asdf", "asdf"}, "_"))
	fmt.Println("Split", s.Split("h__2_2", "_"))

	print(s.Repeat("a", 5))
	fmt.Println("replace", s.Replace("foo", "o", "1", 1))
	fmt.Println("replace", s.Replace("foo", "o", "1", 2))
}
