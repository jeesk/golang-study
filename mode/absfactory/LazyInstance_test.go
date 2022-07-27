package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, getLazyInstance(), getLazyInstance())
}
func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if getLazyInstance() != getInstance() {
				b.Errorf("test faild ")
			}
		}
	})
}
