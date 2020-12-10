package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	b64 := os.Args[1]
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		b, err = base64.URLEncoding.DecodeString(b64)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(b))
}
