package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	encodedString := base64.StdEncoding.EncodeToString([]byte(os.Args[1]))
	fmt.Printf("StdEncoding:     %s\n", encodedString)
	encodedString = base64.RawStdEncoding.EncodeToString([]byte(os.Args[1]))
	fmt.Printf("Raw StdEncoding: %s\n", encodedString)
	encodedString = base64.RawURLEncoding.EncodeToString([]byte(os.Args[1]))
	fmt.Printf("URL StdEncoding: %s\n", encodedString)
}
