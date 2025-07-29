package basic

func ExecNumber(num int) int {
	number := num
	if number == 2 {
		return num * 3
	} else if num == 3 {
		return num * 4
	} else {
		return num //jika semua kondisi tidak ada yang memenuhi maka balikan inputnya
	}
}

func ValidateNumberAvailable(num int) string {
	if ExecNumber(num) > num {
		return "tersedia"
	} else {
		return "tidak tersedia"
	}
}
