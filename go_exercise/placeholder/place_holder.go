package main

import "fmt"

func main() {
	str1 := "hava"
	str2 := "çok"
	str3 := "güzel"
	aNumber := 45
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)

	fmt.Println("string length:", stringLength)

	fmt.Printf("value of aNumber: %v\n", aNumber)
	fmt.Printf("value of isTrue: %v\n", isTrue)
	fmt.Printf("value of aNumber as float: %.2f\n", float64(aNumber))
	fmt.Printf("data types: %T , %T,%T,%T and %T\n", str1, str2, str3, aNumber, isTrue)

	my_String := fmt.Sprintf("data types: %T , %T,%T,%T and %T", str1, str2, str3, aNumber, isTrue)

	fmt.Printf(my_String)

}
