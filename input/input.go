package input

type Type int

const (
	Text Type = iota
	Password
	Date
)

type Input struct {
	inputType string
	name string
}

func NewInput(inputType string, name string) Input {
	return Input {
		inputType: inputType,
		name: name,
	}
}
