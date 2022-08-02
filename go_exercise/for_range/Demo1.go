package for_range

import "fmt"

var i = []int{1, 2, 3, 4, 5}

func for_range() {
	for index, value := range i {
		fmt.Printf("index: %d value:%d\n", index, value)
	}

	//map

	baskentler := map[string]string{"TURKEY": "ISTANBUL", "FRANCE": "PARIS", "GERMANY": "BERLIN"}
	for i, value := range baskentler {
		fmt.Println("map item capital of :", i, "is", value)
	}
}

func Demo2() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	toplam := 0
	for _, rakam := range sayilar {
		if rakam%2 != 0 {
			toplam = toplam + rakam
		}
	}
	fmt.Println(toplam)
}
