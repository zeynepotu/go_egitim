package functions

func Topla(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result

}

func ToplaVariadic(params ...int) int {
	result := 0
	for i := 0; i < len(params); i++ {
		result = result + params[i]
	}
	return result
}
