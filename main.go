package main
import (
	"github.com/chrisvire/goreleaser-example/bar"
)

func main() {
	println("Ba dum, tss!")
	f,_ := NewFoo("bar")
	f.Print()

	b,_ := bar.NewBar("baz")
	b.Print()
}
