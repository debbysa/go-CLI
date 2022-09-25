package calculator

// penjumlahan
type Calculator interface {
	Penjumlahan() int
	Pengurangan() int
}

type dataNum struct {
	num1 int
	num2 int
}

func NewDataNum(num1 int, num2 int) Calculator {
	return &dataNum{num1: num1, num2: num2}
}

func (a dataNum) Penjumlahan() int {
	return a.num1 + a.num2
}

func (a dataNum) Pengurangan() int {
	if a.num2 > a.num1 {
		return a.num2 - a.num1
	}
	return a.num1 - a.num2
}
