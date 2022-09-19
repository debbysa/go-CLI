package main

import "fmt"

// penjumlahan
type calculator interface {
	penjumlahan() int
	pengurangan() int
}

type dataNum struct {
	num1 int
	num2 int
}

func (a dataNum) penjumlahan() int {
	return a.num1 + a.num2
}

func (a dataNum) pengurangan() int {
	if a.num2 > a.num1 {
		return a.num2 - a.num1
	}
	return a.num1 - a.num2
}

// interface luas persegi panjang
type perhitungan interface {
	luasPersegiPanjang() int
}
type dataPersegiPanjang struct {
	p int
	l int
}

func (d dataPersegiPanjang) luasPersegiPanjang() int {
	return d.p * d.l
}

// interface luas persegi
type perhitunganPersegi interface {
	luasPersegi() int
}

type dataPersegi struct {
	sisi int
}

func (d dataPersegi) luasPersegi() int {
	return d.sisi * d.sisi
}

func main() {
	fmt.Println("---Menu---")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Luas Bangun")
	var choice int

	fmt.Print("Masukkan pilihan anda: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println(err)
	}

	switch choice {
	case 1:
		fmt.Println("=== Penjumlahan ===")
		var calculatorResult calculator

		var num1 int
		var num2 int

		fmt.Print("Masukkan angka pertama: ")
		_, err = fmt.Scanf("%d", &num1)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("Masukkan angka kedua: ")
		_, err = fmt.Scanf("%d", &num2)
		if err != nil {
			fmt.Println(err)
		}

		calculatorResult = dataNum{
			num1,
			num2,
		}

		fmt.Println("Penjumlahan (num1 + num2) = ", calculatorResult.penjumlahan())
		break
	case 2:
		fmt.Println("=== Pengurangan ===")
		var calculatorResult calculator

		var num1 int
		var num2 int

		fmt.Print("Masukkan angka pertama: ")
		_, err = fmt.Scanf("%d", &num1)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Print("Masukkan angka kedua: ")
		_, err = fmt.Scanf("%d", &num2)
		if err != nil {
			fmt.Println(err)
		}

		calculatorResult = dataNum{
			num1,
			num2,
		}

		fmt.Println("Pengurangan (num1 - num2 / num2 - num1) = ", calculatorResult.pengurangan())
		break
	case 3:
		fmt.Println("Menghitung luas bangun")
		fmt.Println("1. Persegi Panjang")
		fmt.Println("2. Persegi")

		fmt.Print("Masukkan pilihan anda: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
		}

		switch choice {
		case 1:
			var p int
			var l int

			fmt.Print("Masukkan panjang: ")
			_, err = fmt.Scan(&p)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("Masukkan lebar: ")
			_, err = fmt.Scan(&l)
			if err != nil {
				fmt.Println(err)
			}

			var perhitunganResult perhitungan

			perhitunganResult = dataPersegiPanjang{
				p, l,
			}
			fmt.Println("Luas persegi panjang (p x l) = ", perhitunganResult.luasPersegiPanjang())
		case 2:
			var sisi int
			fmt.Print("Masukkan sisi: ")
			_, err = fmt.Scan(&sisi)
			if err != nil {
				fmt.Println(err)
			}

			var perhitunganResult perhitunganPersegi

			perhitunganResult = dataPersegi{
				sisi,
			}
			fmt.Println("Luas Persegi (s x s) = ", perhitunganResult.luasPersegi())
			break
		default:
			fmt.Println("pilihan tidak ada")
		}

		break

	default:
		fmt.Println("pilihan tidak ada")
	}

}
