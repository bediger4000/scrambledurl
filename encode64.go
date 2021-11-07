package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	encodedString := base64.StdEncoding.EncodeToString([]byte(os.Args[1]))
	fmt.Println(encodedString)
}
