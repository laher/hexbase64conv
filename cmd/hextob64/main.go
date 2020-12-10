package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	h := os.Args[1]
	b, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}

	fmt.Println("url: ", base64.URLEncoding.EncodeToString(b))
	fmt.Println("raw-url: ", base64.RawURLEncoding.EncodeToString(b))
	fmt.Println("raw-std:", base64.RawStdEncoding.EncodeToString(b))
	fmt.Println("std:", base64.StdEncoding.EncodeToString(b))
}
