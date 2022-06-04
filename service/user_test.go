package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/conn"
	"testing"
)

func BenchmarkCheckUserNameRepeat(b *testing.B) {
	conn.InitGorm()
	fmt.Println(ContainsName("tom"))
}
