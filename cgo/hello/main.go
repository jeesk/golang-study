package main

/*#include <stdio.h>*/
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
