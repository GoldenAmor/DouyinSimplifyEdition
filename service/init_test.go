package service

import "testing"

func BenchmarkTestInit(b *testing.B) {
	Init()
}
