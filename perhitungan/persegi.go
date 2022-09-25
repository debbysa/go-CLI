package perhitungan

// interface luas persegi
type PerhitunganPersegi interface {
	LuasPersegi() int
}

type dataPersegi struct {
	sisi int
}

func NewDataPersegi(sisi int) PerhitunganPersegi {
	return &dataPersegi{sisi: sisi}
}

func (d dataPersegi) LuasPersegi() int {
	return d.sisi * d.sisi
}
