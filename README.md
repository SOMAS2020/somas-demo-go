## What am I seeing?

A slightly more fun MWE for `golang`!

Goroutines are used two times here! I'll go into those in a bit...

## Trying this out

### Requirements
- Golang 1.15.x

### Build, run, profit!
- `go build`
- `./somas-demo-go`

## Interpreting the output!

Here's an example:

```
2020/11/18 00:25:16 CALLFUNC2
2020/11/18 00:25:16 CALLFUNC1
2020/11/18 00:25:16 A: Received '{69 420}'
2020/11/18 00:25:16 Got {69 420} from client A
2020/11/18 00:25:16 B: Received '{69 420}'
2020/11/18 00:25:16 Got {69 420} from client B
2020/11/18 00:25:16 CALLFUNC2
2020/11/18 00:25:16 A: Received '{69 420}'
2020/11/18 00:25:16 Got {69 420} from client A
2020/11/18 00:25:16 B: Received '{69 420}'
2020/11/18 00:25:16 Got {69 420} from client B
2020/11/18 00:25:16 CALLFUNC1
2020/11/18 00:25:16 A: Received '{420 420}'
2020/11/18 00:25:16 Got {420 420} from client A
2020/11/18 00:25:16 B: Received '{420 420}'
2020/11/18 00:25:16 Got {420 420} from client B
2020/11/18 00:25:16 A: Received '{420 420}'
2020/11/18 00:25:16 Got {420 420} from client A
2020/11/18 00:25:16 B: Received '{420 420}'
2020/11/18 00:25:16 Got {420 420} from client B
```

(It's different every time!)

> It is left as an exercise to the reader as to why this happens...

### The second instance of goroutines
(This one is easier to explain)
- See `./internal/server/server.go`
- In `CallFunc1()`, I have a comment there explaining what's happening, as for channels, see the next (first) instance...

### The first instance of goroutines
- See `main.go`.
- Basically, `go <function>` spins of a lightweight concurrent thread called a goroutine (warning: not necessarily parallel/multithreaded!)
- `<-` arrow basically waits until that channel (or `chan`) is filled

Again, if you don't understand, don't worry. Go (ha ha ha) have a look at [A Tour Of Go](https://tour.golang.org/)!

### Thoughts
- `go` is a sweet language. Just look at the tooling required. Just the `go` binary in your system!
- Concurrency is *FREE*
- There definitely is overhead for learning, but [A Tour Of Go](https://tour.golang.org/) is a good brief introduction most people can finish with ~3 days provided a sufficient base in C++/Python.
- Free types!
- Free pedantic language server so no one commits even wrong whitespace!
- Honestly, the hard part is the infra setup. Once this is set up, look at whatever code you need to write in `./internal/clients/clienta/client`. It becomes trivial and you can be productive in implementing the SOMASity.