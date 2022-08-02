package conditionals

import "fmt"

func Demo1() {
	//3 adet int değişken tanımla Ekrana en büyük olanı yazınız.
	var1 := 150
	var2 := 250
	var3 := 350

	if var1 > var2 && var1 > var3 {
		fmt.Println(var1)
	}
	if var2 > var1 && var2 > var3 {
		fmt.Println(var2)
	}
	if var3 > var1 && var3 > var2 {
		fmt.Println(var3)
	}

}