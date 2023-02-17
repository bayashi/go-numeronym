package numeronym

import (
	"regexp"
	"strconv"
)

func Numeronize(word string) string {
	length := len(word)
	if length >= 4 {
		if matched, _ := regexp.Match(`^[a-zA-Z]+$`, []byte(word)); matched {
			return word[:1] + strconv.Itoa(length-2) + word[length-1:]
		}
	}

	return word
}
