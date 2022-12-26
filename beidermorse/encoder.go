package beidermorse

type Encoder struct{}

func NewEncoder() *Encoder {
	return &Encoder{}
}

func (e *Encoder) Encode(input string) []string {
	return nil // @todo
}
