package clienta

import (
	"fmt"
	"log"

	"github.com/SOMAS2020/somas-demo-go/internal/common"
)

func init() {
	common.RegisterClient(&Client{"A"})
}

type Client struct {
	id string
}

func (c *Client) GetId() (s string) {
	s = c.id
	return
}

func (c *Client) Func1(arg common.Func1Arg, ch chan common.Func1Ret) {
	c.Log("Received '%v'", arg)
	ret := common.Func1Ret{
		MemesStr: arg.MemesStr,
		MemesInt: arg.MemesInt,
	}
	ch <- ret
}

func (c *Client) Func2(arg common.Func2Arg, ch chan common.Func2Ret) {
	c.Log("Received '%v'", arg)
	ret := common.Func2Ret{
		MoreMemesInt:   arg.MoreMemesInt,
		MoreMemesFloat: arg.MoreMemesFloat,
	}
	ch <- ret
}

func (c *Client) Log(format string, a ...interface{}) {
	log.Printf("%v: %v", c.id, fmt.Sprintf(format, a...))
}
