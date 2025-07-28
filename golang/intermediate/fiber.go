package intermediate

import "fmt"

func Rata_rata_nilai() float64 {
	var nilai [5]float64

	fmt.Println("Masukkan 5 nilai:")

	// Input nilai dari user
	for i := 0; i < len(nilai); i++ {
		fmt.Printf("Nilai ke-%d: ", i+1)
		fmt.Scan(&nilai[i])
	}

	// Menampilkan nilai dan menghitung total
	total := 0.0
	fmt.Println("\nNilai yang dimasukkan:")
	for i, val := range nilai {
		fmt.Printf("Siswa %d: %.2f\n", i+1, val)
		total += val
	}

	rataRata := total / float64(len(nilai))
	return rataRata
}

func Total_nilai() float64 {
	var nilai [5]float64

	fmt.Println("Masukkan 5 nilai:")

	// Input nilai dari user
	for i := 0; i < len(nilai); i++ {
		fmt.Printf("Nilai ke-%d: ", i+1)
		fmt.Scan(&nilai[i])
	}

	// Menampilkan nilai dan menghitung total
	total := 0.0
	fmt.Println("\nNilai yang dimasukkan:")
	for i, val := range nilai {
		fmt.Printf("Siswa %d: %.2f\n", i+1, val)
		total += val
	}
	return total
}
