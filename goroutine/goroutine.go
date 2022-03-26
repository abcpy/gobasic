package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	/*for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello form "+"goroutine %d\n", i)
			}
		}(i)
	}

	time.Sleep(time.Millisecond)*/

	/*go hello()
	time.Sleep(time.Second)*/

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main function")

}
