package chap8

import (
	"fmt"
	"testing"
)

func TestGetPath(t *testing.T) {
	b := make([][]SQ, 8)
	for i := range b {
		b[i] = make([]SQ, 8)
	}
	b[7][0].allowed = true
	path := make([]Position, 1)
	s := Position{0, 0}
	GetPath(b, s, &path)
	fmt.Println(path)
}
