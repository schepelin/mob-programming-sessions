package gameoflife

// Life contains current and next generation of tha game
type Life struct {
	width        uint
	height       uint
	currentState [][]bool
	nextState    [][]bool
}

// NewLife creates new life with an initial state
func NewLife(initial [][]bool) *Life {
	l := &Life{
		height:       uint(len(initial)),
		width:        uint(len(initial[0])),
		currentState: initial,
	}
	return l
}
