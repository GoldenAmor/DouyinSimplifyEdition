package repository

import "testing"

func BenchmarkTestInit(b *testing.B) {
	Init()
}
