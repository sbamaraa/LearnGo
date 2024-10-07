package print_foobar

type FooBar struct {
	n int
	ch chan bool
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n: n,
		ch: make(chan bool),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.ch <- true
		printFoo()
		fb.ch <- true
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch
		<-fb.ch
		printBar()
	}
}
