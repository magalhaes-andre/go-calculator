package calculator

type Operation struct {
	First float64
	Second float64
	TypeOf string
	Result float64
}

func NewOperation(first float64, second float64, typeOf string) Operation{
	return Operation{first, second, typeOf, 0}
}