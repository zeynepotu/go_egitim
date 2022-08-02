package slices

import "fmt"

func Demo1() {
	sehirler := make([]string, 3)
	fmt.Println(sehirler)
	sehirler[0] = "Cagla"
	sehirler[1] = "Leyla"
	sehirler[2] = "Ugursay"
	fmt.Println(sehirler)
	sehirler = append(sehirler, "Ali")
	sehirlerKopya := make([]string, len(sehirler))
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	fmt.Println(sehirler[1:3])

}
