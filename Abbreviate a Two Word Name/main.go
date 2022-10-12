package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) (abb string) {

	//my solution was the best

	splitName := strings.Split(name, " ")
	abb = strings.ToUpper(splitName[0][0:1] + "." + splitName[1][0:1])
	return
}

func main() {
	fmt.Println(AbbrevName("oleg Orlovsky"))

}
