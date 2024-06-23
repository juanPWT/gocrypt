package gocrypt

import (
	"errors"

	"github.com/juanPWT/gocrypt/algoritm"
)

func Decode(cyper, key, algo string) (string, error) {
	var result string
	switch algo {
	case Mufaseh:
       result = algoritm.MufasehDecoding(cyper, key) 
	default:
		return "", errors.New("algoritm must be needed")
	}

    return result, nil
}
