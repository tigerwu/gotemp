package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {

}

func (iPhone ApplePhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// func main() {
// 	var nphone Phone
// 	var iphone Phone

// 	fmt.Printf("hello, Call me now\n")

// 	nphone = new(NokiaPhone)
// 	nphone.call()

// 	iphone = new(ApplePhone)
// 	iphone.call()
// }