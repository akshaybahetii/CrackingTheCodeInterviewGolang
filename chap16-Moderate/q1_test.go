package chap16

import (
	"fmt"
	"testing"
)

func TestExchangeNum(t *testing.T) {
	a := 4
	b := 10
	ExchangeNum(&a, &b)
	fmt.Println(a, b)
}
