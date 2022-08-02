package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng"))
	fmt.Println(s.HasSuffix(isim, "en"))
	fmt.Println(s.Index(isim, "g"))
	harfler := []string{"e", "n", "g", "i", "n"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)
	fmt.Println(s.Replace(sonuc, "*", "+", -1))
	fmt.Println(s.Split(sonuc, "*"))

}
