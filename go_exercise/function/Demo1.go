package functions

func Islem(sayi1 int, sayi2 int) (int, int, int, float64) {
	toplama := sayi1 + sayi2
	cikarma := sayi1 - sayi2
	carpma := sayi1 * sayi2
	bolme := float64(sayi1 / sayi2)

	return toplama, cikarma, carpma, bolme
}
