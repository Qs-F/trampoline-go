package trampoline

type F func() F

type Arg interface{}

type Any interface{}

type Base struct {
	ret chan Any
}

func (b *Base) Done() F {
	return b.Done
}

func (b *Base) Return(result Any) {
	b.ret <- result
}

var B Base = Base{
	ret: make(chan Any, 1),
}

func Trampoline(f func(args ...Arg) F) func(args ...Arg) Any {
	return func(args ...Arg) Any {
		res := f(args...)
		for {
			select {
			case ret := <-B.ret:
				return ret
			default:
			}
			res = res()
		}
	}
}
