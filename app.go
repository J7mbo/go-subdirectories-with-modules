package main

import (
	"app/bar"
	"app/foo"
	"app/foo/baz"
	"fmt"
)

func main() {
	fmt.Println(foo.Foo())
	fmt.Println(bar.Bar())
	fmt.Println(baz.Baz())
}
