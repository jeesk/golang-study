package main

// 恶汉单利
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func getInstance() *Singleton {
	return singleton
}
