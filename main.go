package main

type F func() F

type Base struct {
	ch chan int
}

func (b *Base) base() F {
	b.ch <- 1
	return b.base
}

var B Base = Base{
	ch: make(chan int, 1),
}

func down(n int) F {
	if n == 1 {
		return B.base
	}
	return func() F { return down(n - 1) }
}

func trampoline(f func(n int) F) func(n int) int {
	return func(arg int) int {
		res := f(arg)
		for {
			select {
			case <-B.ch:
				return 1
			default:
			}
			res = res()
		}
	}
}

func main() {
	// down(100000000)
	// t := trampoline(down)
	// t(100000000)
	for i := 0; i < 100000000; i++ {
	}
}
