package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcSquares2(number int, squareop chan int) {
	for number != 0 {
		digit := number % 10
		squareop <- digit
		number /= 10
	}
	close(squareop)
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	squareop := make(chan int)
	go calcSquares2(number, squareop)
	for digit := range squareop {
		sum += digit * digit
	}
	cubeop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	/*var a chan int
	if a == nil {
		fmt.Println("channel is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}*/

	/*done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done
	fmt.Println("Main received data")*/

	/*	number := 589
		sqrch := make(chan int)
		cubech := make(chan int)
		go calcSquares(number, sqrch)
		go calcCubes(number, cubech)
		squares, cubes := <-sqrch, <-cubech
		fmt.Println("Final output", squares+cubes)*/

	/*ch := make(chan int)
	ch <- 5  //deadlock!*/

	/*sendch := make(chan<- int)
	go sendData(sendch)
	fmt.Println(<-sendch)*/

	/*sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)*/

	/*	ch := make(chan int)
		go producer(ch)*/
	/*for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}*/

	/*	for v := range ch {
		fmt.Println("Received ", v)
	}*/

	number := 589
	cubech := make(chan int)
	go calcCubes2(number, cubech)
	cubes := <-cubech
	fmt.Println("Final output", cubes)

}
