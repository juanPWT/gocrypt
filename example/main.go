package main

import (
	"fmt"

	"github.com/juanPWT/gocrypt"
)

var key = "hello"

func main() {
	plainText := "juan"

	encodeText, err := gocrypt.Encode(plainText, key, gocrypt.Mufaseh)
    if err != nil {
        panic(err)
    }

    fmt.Printf("encode text: %s\n", encodeText)

    decodeText, err := gocrypt.Decode(encodeText, key, gocrypt.Mufaseh)
    if err != nil {
        panic(err)
    }

    fmt.Printf("decode text: %s\n", decodeText)

}
