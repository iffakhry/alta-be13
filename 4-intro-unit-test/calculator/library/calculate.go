package library

var data int

func Penambahan(bil1, bil2 int) int {
	hasil := bil1 + bil2
	return hasil
}

func Pengurangan(bil1, bil2 int) int {
	hasil := bil1 - bil2
	return hasil
}

func Pembagian(bil1, bil2 int) int {
	hasil := bil1 / bil2
	return hasil
}

func KonversiNilai(nilai int) string {
	var result string
	if nilai > 85 && nilai <= 100 {
		result = "A"
	} else if nilai >= 65 && nilai <= 85 {
		result = "B+"
	} else if nilai >= 50 && nilai < 65 {
		result = "B"
	} else if nilai >= 0 && nilai < 50 {
		result = "D"
	} else {
		result = "format nilai tidak sesuai"
	}
	return result
}
