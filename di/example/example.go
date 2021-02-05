package main

import (
	"fmt"

	"github.com/edgexfoundry/go-mod-bootstrap/di"
)

type foo struct {
	FooMessage string
}

func NewFoo(m string) *foo {
	return &foo{
		FooMessage: m,
	}

}

type bar struct {
	BarMessage string
	Foo        *foo
}

func NewBar(m string, foo *foo) *bar {
	return &bar{
		BarMessage: m,
		Foo:        foo,
	}

}

func main() {
    di.	
	container := di.NewContainer(
		di.ServiceConstructorMap{
			"foo": func(get di.Get) interface{} {
				return NewFoo("fooMessage1")
			},
			"bar": func(get di.Get) interface{} {
				return NewBar("barMessage", get("foo").(*foo))
			},
		})

	a := container.Get("foo").(*foo)
	fmt.Println(a.FooMessage)
	b := container.Get("bar").(*bar)

	fmt.Println(b.BarMessage)
	fmt.Println(b.Foo.FooMessage)
}
