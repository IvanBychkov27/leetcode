package main

import (
	"fmt"
	"testing"
)

func Test_noReplay(t *testing.T) {
	d := [][]byte{{'('}, {')'}, {'('}, {')'}}
	fmt.Println(d)

	res := noReplay(d)
	fmt.Println(res)
}

func TestCheck(t *testing.T) {
	data := []byte("((())(()))")
	res := check(data)
	fmt.Println("res =", res)
}
