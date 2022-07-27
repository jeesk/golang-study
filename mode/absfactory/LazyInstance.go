package main

import "sync"

var (
	lazySingleton *Singleton
	once          = sync.Once{}
)

func getLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
