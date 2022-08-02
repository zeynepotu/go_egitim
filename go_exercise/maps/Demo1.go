package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"
	fmt.Println(sozluk["table"])
	delete(sozluk, "book")
	fmt.Println(sozluk)

	deger, varMi := sozluk["pencil"]
	fmt.Println(deger, varMi)

	sozluk2 := map[string]string{"table": "Masa", "book": "Kitap"}
	fmt.Println(sozluk2)

}
