package main

import (
	"time"
)

func server1(ch chan string) {
	//time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	//time.Sleep(3 * time.Second)
	ch <- "from server2"

}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	/*out1 := make(chan string)
	out2 := make(chan string)
	go server1(out1)
	go server2(out2)
	select {
	case s1 := <-out1:
		fmt.Println(s1)
	case s2 := <-out2:
		fmt.Println(s2)
	}*/

	/*ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}*/

	/*	ch := make(chan string)
		select {
		case <-ch:
		default:
			fmt.Println("default case executed")
		}*/
	//select {}
}
