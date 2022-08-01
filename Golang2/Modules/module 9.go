package Modules

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Module9() {
	fmt.Println("module 9")
	data, err := ioutil.ReadFile("/Users/Alan/Desktop/Golang Week 5.txt")

	// feedback message in case of error
	fmt.Println(err)

	// convert the file contents to a string and display it
	doc := string(data)
	words := strings.Split(doc, " ")
	for _, ii := range words {
		fmt.Println(string(ii))
	}
	//fmt.Print(string(data))
}

func Test() { //not an activity
	fmt.Println("Test")
	test := "/hello:: sasdfesdfe"
	splitString := strings.Split(test, "::")
	fmt.Println(splitString)
	str := ""
	str = test[0:1]
	fmt.Println(str)
	str = test[0:1]
	fmt.Println(splitString[0][1:])
}
