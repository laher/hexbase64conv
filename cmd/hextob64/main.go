package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
)

func main() {
	u := flag.Bool("u", false, "url-style encoding")
	s := flag.Bool("s", false, "standard encoding")
	flag.Parse()
	h := flag.Args()[0]
	b, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	if *u {
		fmt.Println(base64.URLEncoding.EncodeToString(b))
		return
	}
	if *s {
		fmt.Println(base64.StdEncoding.EncodeToString(b))
		return
	}
	fmt.Println("url: ", base64.URLEncoding.EncodeToString(b))
	fmt.Println("std: ", base64.StdEncoding.EncodeToString(b))
}
