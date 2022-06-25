package main

import (
	"fmt"
	"unsafe"
)

func main() {
	stu := &Stu{
		age: 1111,
	}
	stu.Pkg = &Pkg{}

	pkg := Pkg{
		nameL: "asdf",
	}
	sizeof := unsafe.Sizeof(pkg)
	stu.okg = pkg

	fmt.Println(sizeof)
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(stu)) + sizeof + 1*unsafe.Sizeof(int(0))))
	fmt.Println(*age)
}

type Stu struct {
	*Pkg
	okg Pkg
	//Pkg
	age int
}

type Pkg struct {
	nameL string
}
