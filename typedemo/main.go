package main

type (
	Error struct {
		*string
		name string
		code int32
	}
)

func main() {
	sp := "asdf"
	e := &Error{string: &sp, name: "hello", code: 12}
	println(e)
}
