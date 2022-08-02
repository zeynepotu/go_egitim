package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	isim := "Zeynep"
	fmt.Println(s.Count(isim, "e"))
	fmt.Println(s.Contains(isim, "a"))
	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim))
	fmt.Println(s.ToUpper(isim))

}
