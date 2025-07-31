package main

import (
	"fmt"
	"golang/basic"
	"golang/intermediate"
)

func ValidateNumberAvailable(num int) string {
	if basic.CheckNumber(num) > num {
		return "tersedia"
	} else {
		return "tidak tersedia"
	}
}

func main() {
	var num int
	//array dulu
	basic.Array(intermediate.Rata_rata_nilai(), intermediate.Total_nilai())
	//condition
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num) //input dari user

	result := ValidateNumberAvailable(num)
	message := fmt.Sprintf("Angka %d, %s", basic.CheckNumber(num), result)
	fmt.Println(message)
}
