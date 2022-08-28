package main

type resp struct {
	Page   int
	Fruits []string
}
type resp1 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// golang 的json
}
