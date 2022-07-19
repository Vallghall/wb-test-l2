package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	res, err := unpack("a0")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// unpack func unpacks strings given in the right format
func unpack(s string) (string, error) {
	// return empty string right ahead
	if s == "" {
		return "", nil
	}
	// check if the string follows the needed format. Returns an error if it doesn't
	if ok, _ := regexp.MatchString(`^([0-9])|[^\\][0-9]{2}|[^\\]\\[a-z]`, s); ok {
		return "", errors.New("invalid string")
	}
	// init Builder
	sb := strings.Builder{}
	// init flag for watching escape characters
	escaped := false
	// temp stores the previous rune to multiply it with number
	var temp rune = 0
	for _, char := range s {
		if unicode.IsLetter(char) || escaped {
			sb.WriteRune(char) // accumulating string within sb
			temp = char        // set current temp
			escaped = false    // reset escape flag
			continue
		}
		// set escaped flag to true to escape the next character
		if char == '\\' {
			escaped = true
			continue
		}
		// multiplication
		for i := 0; i < int(char-'1'); i++ {
			sb.WriteRune(temp)
		}
	}
	// return accumulated string
	return sb.String(), nil
}
