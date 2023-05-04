package main

// package main

import (
	"fmt"
	"io/ioutil"
	"junit_parse/util"
	"os"
)

func main() {
	xmlFile, err := os.Open("junit.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened junit.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(xmlFile)
	fmt.Println("Bytes:", string(byteValue))
	if err != nil {
		fmt.Println(err)
	}

	result, err := util.ConvertResultJson(byteValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}