package bar

type Bar struct {
	message		string
}

func NewBar(message string) (*Bar, error) {
	b := &Bar{
		message: message,
	}
	return b, nil
}

func (b *Bar) Print() error {
	println(b.message)
	return nil
}
