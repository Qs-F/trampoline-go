package main

import (
	"fmt"

	"github.com/Qs-F/trampoline-go"
)

func down(args ...trampoline.Arg) trampoline.F {
	n := args[0].(int)
	fmt.Println(n)
	if n == 1 {
		trampoline.B.Return(1)
		return trampoline.B.Done
	}
	return func() trampoline.F {
		return down(n - 1)
	}
}

func fib(args ...trampoline.Arg) trampoline.F {
	count := args[0].(int)
	a := args[1].(int)
	b := args[2].(int)
	fmt.Println(count, a)
	if count == 1 {
		trampoline.B.Return(a)
		return trampoline.B.Done
	}
	return func() trampoline.F {
		return fib(count-1, b, a+b)
	}
}

func main() {
	t1 := trampoline.Trampoline(down)
	t1(10)

	t2 := trampoline.Trampoline(fib)
	t2(10, 1, 1)
}
