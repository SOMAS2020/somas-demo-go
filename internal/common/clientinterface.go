package common

type Func1Arg struct {
	MemesStr string
	MemesInt int
}

type Func1Ret struct {
	MemesStr string
	MemesInt int
}

type Func2Arg struct {
	MoreMemesInt   int
	MoreMemesFloat float32
}

type Func2Ret struct {
	MoreMemesInt   int
	MoreMemesFloat float32
}

type Client interface {
	Func1(Func1Arg, chan Func1Ret)
	Func2(Func2Arg, chan Func2Ret)
	Log(format string, a ...interface{})
	GetId() string
}
