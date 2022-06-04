package service

import "testing"

func BenchmarkCDN(b *testing.B) {
	b.ResetTimer()
	Upload("1.mp4")
}
