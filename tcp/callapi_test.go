package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go startAPI()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestSumLen(t *testing.T) {
	ch := sumLen("!@#$%^&*()")
	fmt.Println(<-ch)
}

func TestReset(t *testing.T) {
	ch := reset()
	fmt.Println(<-ch)
}

func BenchmarkSumLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := sumLen("1")
		res := <-ch
		if res != "Too Many Requests\n" {
			fmt.Println(res)
		}
	}
}
