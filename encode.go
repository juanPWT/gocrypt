package gocrypt

import (
	"errors"

	"github.com/juanPWT/gocrypt/algoritm"
)

var (
    Mufaseh = "mufaseh"
)


func Encode(plain, key, algo string) (string, error) {
    var result string
	switch algo {
	case Mufaseh:
		result = algoritm.MufasehEncoding(plain, key)
    default: 
        return "", errors.New("encoding algoritm must be needed")
	}

    return result, nil
}
