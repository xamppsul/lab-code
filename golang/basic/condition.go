package basic

func CheckNumber(num int) int {
	number := num
	if number == 2 {
		return num * 3
	} else if num == 3 {
		return num * 4
	} else {
		return num //jika semua kondisi tidak ada yang memenuhi maka balikan inputnya
	}
}
