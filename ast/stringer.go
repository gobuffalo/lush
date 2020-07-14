package ast

// LushStringer is implemented by any value that has a LushString method, which defines the Lush syntax for that value. The LushString method is used to print values passed as an operand to a %*v format.
type LushStringer interface {
	LushString() string
}
