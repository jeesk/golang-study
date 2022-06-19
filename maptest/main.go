package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string)
	m["hello"] = "1234"
	testMap(m)
	fmt.Println(m)
	var syncMp sync.Map
	syncMp.Store("hello", "1234")
	testSyncMap(&syncMp)
	load, _ := syncMp.Load("hello")
	fmt.Println(load)
}

func testMap(mp map[string]string) {
	mp["hello"] = "hello"
}

func testSyncMap(mp *sync.Map) {
	mp.Store("hello", "hello")
}
