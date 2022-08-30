// print N numbers using two threads one for even and 1 for odd
package main

import (
	"fmt"
)

type handler struct {
	state int
}

func (n *handler) handleOdd(ch chan int, state int) {

	ch <- state

}

func (n *handler) handleEven(ch chan int, state int) {

	ch <- state

}

func main() {
	N := 10
	ch := make(chan int, 1)
	for i := 1; i <= N; i++ {

		if i%2 == 0 {
			handler := &handler{}
			go handler.handleEven(ch, i)
			val := <-ch
			fmt.Println(val)
		}
		if i%2 != 0 {
			handler := &handler{}
			go handler.handleOdd(ch, i)
			val := <-ch
			fmt.Println(val)
		}
	}
}
