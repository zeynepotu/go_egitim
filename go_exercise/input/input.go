package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Demo1() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("enter a number:")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value of number:", f)
	}
}
