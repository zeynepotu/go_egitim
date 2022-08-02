package enum

import "fmt"

type Brand string

const (
	FACEBOOK  Brand = "facebook"
	MICROSOFT Brand = "microsoft"
	GOOGLE    Brand = "google"
	DIJIBEL   Brand = "dijibel"
)

func PrintBrand(brand Brand) {
	fmt.Println(brand)
}
