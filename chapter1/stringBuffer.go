package main

import (
	"log"
	"strings"
)

func StringConcat() {
	var builder strings.Builder
	src := []string{"Back", "To", "The", "Future", "Part", "III"}
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())
}
