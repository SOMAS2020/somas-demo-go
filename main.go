package main

import (
	"github.com/SOMAS2020/somas-demo-go/internal/server"
)

func callFunc1(s server.Server, c chan error) {
	s.CallFunc1()
	c <- nil
}

func callFunc2(s server.Server, c chan error) {
	s.CallFunc2()
	c <- nil
}

func main() {
	s := server.SOMASServerFactory()

	// Let's spin off some goroutines!!
	c1 := make(chan error)
	c2 := make(chan error)
	c3 := make(chan error)
	c4 := make(chan error)
	go callFunc1(s, c1)
	go callFunc2(s, c2)
	go callFunc1(s, c3)
	go callFunc2(s, c4)

	// and we wait for them to complete :)
	<-c1
	<-c2
	<-c3
	<-c4

	// of course this can be done more elegantly!
}
