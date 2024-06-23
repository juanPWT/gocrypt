package gocrypt

import (
	"testing"
)

var key = "hello" 

func TestMufasehEncoding(t *testing.T) {
	plainText := "juan"
	expectedText := "nbvj"

	cyper, err := Encode(plainText, key, Mufaseh)

	if cyper != expectedText || err != nil {
        t.Fatal("skill issue: mufaseh encode")
	}
}

func TestMufasehDecoding(t *testing.T) {
    cipherText := "nbvj"
    expectedText  := "juan"

    decodeText, err := Decode(cipherText, key, Mufaseh)

    if decodeText != expectedText || err != nil {
        t.Fatal("skill issue: mufaseh decode")
    }
}
