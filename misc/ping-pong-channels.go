package main

import (
	"fmt"
)

func play(s string, ball chan int, q chan int) {
	for {
		var hit int
		hit = <-ball
		hit++
		fmt.Println(s, hit)
		if hit == 10 {
			q <- 1
		}
		ball <- hit
	}
}
func main() {
	ball := make(chan int)
	close := make(chan int)
	go play("ping", ball, close)
	go play("pong", ball, close)

	ball <- 1

	<-close

}
