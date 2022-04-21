package main

type Foo struct {
	message		string
}

func NewFoo(message string) (*Foo, error) {
	f := &Foo{
		message: message,
	}
	return f, nil
}

func (f *Foo) Print() error {
	println(f.message)
	return nil
}
