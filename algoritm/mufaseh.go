package algoritm

import "strings"


func MufasehEncoding(p, key string) string {
	// custom alphabet 
    alphabet := costumAlphabet(key)
    reverseAlphabert := reverse(alphabet)
    
    // plain 
    plain := strings.ToLower(p)
    
    var result strings.Builder
    for _, char := range plain {
        index := strings.IndexRune(alphabet, char)
        if index != -1 {
            result.WriteByte(reverseAlphabert[index])
        } else {
            result.WriteRune(char)
        }
    }

    return result.String()
}

func MufasehDecoding(c, key string) string {
    // custom alphabet 
    alphabet := costumAlphabet(key)
    reverseAlphabert := reverse(alphabet)
    
    // cipher
    cipher := strings.ToLower(c) 
    
    var result strings.Builder
    for _, char := range cipher {
        index := strings.IndexRune(reverseAlphabert, char)
        if index != -1 {
            result.WriteByte(alphabet[index])
        } else {
            result.WriteRune(char)
        }
    }

    return result.String()
}


func reverse(s string) string {
    runes := []rune(s)

    for i,j := 0, len(runes)-1; i < j; i,j = i+1,j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }  

    return string(runes)
}

func costumAlphabet(key string) string {
    var alphabet = "abcdefghijklmnopqrstuvwxyz"
    keyLowerCase := strings.ToLower(key) + alphabet
    
    custom := removeDuplicateChar(keyLowerCase)

    return custom
}

func removeDuplicateChar(plain string) string {
    seen := make(map[rune]bool)
    var strBuilder strings.Builder

    for _, char := range plain {
        if !seen[char] {
            seen[char] = true
            strBuilder.WriteRune(char)
        }
    }

    return strBuilder.String()
}
