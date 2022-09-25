package main

import (
	"fmt"
	calc "github.com/debbysa/go-CLI/calculator"
	perhitungan "github.com/debbysa/go-CLI/perhitungan"
)

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
		var calculatorResult calc.Calculator

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

		calculatorResult = calc.NewDataNum(num1, num2)

		fmt.Println("Penjumlahan (num1 + num2) = ", calculatorResult.Penjumlahan())
		break
	case 2:
		fmt.Println("=== Pengurangan ===")
		var calculatorResult calc.Calculator

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

		calculatorResult = calc.NewDataNum(num1, num2)

		fmt.Println("Pengurangan (num1 - num2 / num2 - num1) = ", calculatorResult.Pengurangan())
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

			var perhitunganResult perhitungan.Perhitungan

			perhitunganResult = perhitungan.NewDataPersegiPanjang(p, l)
			fmt.Println("Luas persegi panjang (p x l) = ", perhitunganResult.LuasPersegiPanjang())
		case 2:
			var sisi int
			fmt.Print("Masukkan sisi: ")
			_, err = fmt.Scan(&sisi)
			if err != nil {
				fmt.Println(err)
			}

			var perhitunganResult perhitungan.PerhitunganPersegi

			perhitunganResult = perhitungan.NewDataPersegi(sisi)

			fmt.Println("Luas Persegi (s x s) = ", perhitunganResult.LuasPersegi())
			break
		default:
			fmt.Println("pilihan tidak ada")
		}

		break

	default:
		fmt.Println("pilihan tidak ada")
	}

}
