package main

import (
	"golesson/channels"
	"golesson/string_functions"
)

func main() {

	//conditionals.Demo1()
	//enum.PrintBrand(enum.GOOGLE)
	//slices.Demo1()
	//numTerms, sum := functions.Topla(3, 5, 6)
	//fmt.Println("added:", numTerms, "terms for a total of", sum)
	//sayilar := []int{4, 5, 6}
	//sonuc := functions.ToplaVariadic(sayilar...)
	//fmt.Println(sonuc)
	//for_range.Demo2()

	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)

	//ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım :", carpim)
	//interfaces.Demo1()
	//interfaces.Demo2()
	string_functions.Demo1()
	string_functions.Demo2()
}
