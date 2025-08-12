package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

func fahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan suhu (dengan satuan C atau F): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		fmt.Println("Input tidak valid. Gunakan format: [angka] [C/F]")
		return
	}

	temp, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Input suhu tidak valid.")
		return
	}

	unit := strings.ToUpper(parts[1])
	if unit != "C" && unit != "F" {
		fmt.Println("Satuan suhu tidak valid (gunakan C atau F).")
		return
	}

	var result float64
	if unit == "C" {
		result = celciusToFahrenheit(temp)
		fmt.Printf("%.2f Celcius = %.2f Fahrenheit\n", temp, result)
	} else {
		result = fahrenheitToCelcius(temp)
		fmt.Printf("%.2f Fahrenheit = %.2f Celcius\n", temp, result)
	}
}
