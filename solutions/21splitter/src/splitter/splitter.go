package splitter

import (
	"fmt"
	"strings"
)

func Splitter(msg string) {
	mapper := map[string]int{}
	for index, word := range strings.Split(msg, " ") {
		mapper[word] = index
	}
	for key := range mapper {
		fmt.Println(key)
	}
}
