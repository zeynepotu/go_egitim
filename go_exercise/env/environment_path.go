package main

import (
	"fmt"
	"os"
)

func main() {
	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorAchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	//processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	//processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("KULLANICI ADI:" + uName)
	fmt.Println("KULLANICI ADRESİ:" + uDomain)
	fmt.Println("İŞLEMCİ MİMARİSİ:" + processorAchitecture)
	fmt.Println("GOROOT:" + goRoot)
	fmt.Println("GOPATH:" + goPath)
	fmt.Println("HOMEPATH:" + homePath)

}
