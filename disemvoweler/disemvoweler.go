package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	runes := []rune(strings.Join(os.Args[1:], " "))
	vowels := make([]rune, 10)
	for i := 0; i < len(runes); i++ {
		if strings.ContainsRune("aeiou ", runes[i]) {
			if runes[i] != ' ' {
				vowels = append(vowels, runes[i])
			}
			runes = append(runes[:i], runes[i+1:]...)
			i = 0
		}
	}
	fmt.Printf("%s\n%s\n", string(runes), string(vowels))
}