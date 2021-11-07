package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	originalStringBytes, err := base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		log.Fatalf("Some error occured during base64 decode. Error %s", err.Error())
	}
	fmt.Println(string(originalStringBytes))
}
