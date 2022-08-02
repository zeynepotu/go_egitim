package main

import (
	"fmt"
	"time"
)

func main() {
	//konsol:tarih&saat işlemleri
	//date ile kendi tarih verimizi oluşturuyoruz.

	t := time.Date(2016, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("go launch at %s\n", t.Local())
	fmt.Println("-----------------")
	// time.Now() ile şu anın zaman bilgisini alıyoruz.
	now := time.Now()

	fmt.Printf("the time is:,%s\n", now)
	fmt.Println("-----------------")

	//tarihe 1 gn ekle!
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("tomorrowis %v,%v,%v,%v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())
	fmt.Println("-----------------")

}
