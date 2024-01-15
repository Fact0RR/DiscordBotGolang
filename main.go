package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	//	n our jsonFile
	jsonFile, err := os.Open("conf.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	
	byteValue, _ := io.ReadAll(jsonFile)

	var conf Conf

	json.Unmarshal(byteValue, &conf)

	fmt.Println(conf.Token)
}