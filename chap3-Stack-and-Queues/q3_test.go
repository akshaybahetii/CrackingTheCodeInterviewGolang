package chap3

import (
	"fmt"
	"testing"
)

func TestSetofStacks(t *testing.T) {
	var sss SetofStacks
	ss := &sss
	ss.Push(3)
	ss.Push(32)
	ss.Push(33)
	ss.Push(34)
	ss.Push(35)
	ss.Push(36)
	ss.Push(37)
	ss.Push(38)
	ss.Push(39)
	ss.Push(40)
	ss.Push(41)
	ss.Push(42)
	ss.Push(43)
	ss.Push(44)
	ss.Push(45)
	ss.Push(46)
	fmt.Println(ss)
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss)
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
	fmt.Println(ss.Pop())
}
