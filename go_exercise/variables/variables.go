package main

import "fmt"

func main() {

	/*var message string
	message = "merhaba go"
	var message= "hello world"


	var a, b, c int = 3, 4, 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	*/

	u := 55
	v, n := "abc", true
	var durum bool = false
	fmt.Printf("değişken tipi: %T ", durum)

	fmt.Println(u)
	fmt.Println(v, n)
	var durum2 bool
	var metin1 string = "engin"
	var metin2 string = "ahmet"
	durum2 = metin1 != metin2
	fmt.Println(durum2)
}
