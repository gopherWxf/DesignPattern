package Factory_Method

type number struct {
	a float64
	b float64
}

type Operation interface {
	SetA(float64)
	SetB(float64)
	GetResult() float64
}

func (num *number) SetA(n float64) {
	num.a = n
}
func (num *number) SetB(n float64) {
	num.b = n
}

type addOperation struct {
	number
}

func (num *addOperation) GetResult() float64 {
	return num.a + num.b
}

type subOperation struct {
	number
}

func (num *subOperation) GetResult() float64 {
	return num.a - num.b
}

type mulOperation struct {
	number
}

func (num *mulOperation) GetResult() float64 {
	return num.a * num.b
}

type Factory interface {
	CreateOperation() Operation
}
type AddFactory struct {
}

func (a *AddFactory) CreateOperation() Operation {
	return &addOperation{}
}

type SubFactory struct {
}

func (s *SubFactory) CreateOperation() Operation {
	return &subOperation{}
}

type MulFactory struct {
}

func (m *MulFactory) CreateOperation() Operation {
	return &mulOperation{}
}
