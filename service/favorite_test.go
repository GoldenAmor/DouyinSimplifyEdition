package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conn"
	"testing"
)

func BenchmarkIsFavorite(b *testing.B) {
	conn.InitGorm()
	fmt.Println(IsFavorite(1, 1))
}
