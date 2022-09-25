package perhitungan

// interface luas persegi panjang
type Perhitungan interface {
	LuasPersegiPanjang() int
}
type dataPersegiPanjang struct {
	p int
	l int
}

func NewDataPersegiPanjang(p int, l int) Perhitungan {
	return &dataPersegiPanjang{p, l}
}

func (d *dataPersegiPanjang) LuasPersegiPanjang() int {
	return d.p * d.l
}
