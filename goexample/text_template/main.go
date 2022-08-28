package main

import "text/template"

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value:{{..}}"))
	// golang 的模板类似于js的模板功能一样
}
