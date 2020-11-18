package server

import (
	"log"

	"github.com/SOMAS2020/somas-demo-go/internal/common"

	// so that init runs :)
	_ "github.com/SOMAS2020/somas-demo-go/internal/clients/clienta"
	_ "github.com/SOMAS2020/somas-demo-go/internal/clients/clientb"
)

type Server interface {
	CallFunc1() error
	CallFunc2() error
}

type SOMASServer struct {
	allClients []common.Client
}

func SOMASServerFactory() (s Server) {
	s = &SOMASServer{common.AllClients}
	return
}

func (s *SOMASServer) CallFunc1() (err error) {
	log.Printf("CALLFUNC1")
	for _, client := range s.allClients {
		c := make(chan common.Func1Ret)
		args := common.Func1Arg{
			MemesStr: "420",
			MemesInt: 420,
		}
		// spins off a goroutine! a lightweight concurrent thread :)
		// this is just like async/await here, but we can have a slice of channels
		// and spin off n func calls and wait for all of them to come back :)
		go client.Func1(args, c)
		ret := <-c
		log.Printf("Got %v from client %v", ret, client.GetId())
	}
	return
}

func (s *SOMASServer) CallFunc2() (err error) {
	log.Printf("CALLFUNC2")
	for _, client := range s.allClients {
		c := make(chan common.Func2Ret)
		args := common.Func2Arg{
			MoreMemesInt:   69,
			MoreMemesFloat: 420,
		}
		go client.Func2(args, c)
		ret := <-c
		log.Printf("Got %v from client %v", ret, client.GetId())
	}
	return
}
