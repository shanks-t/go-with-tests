package main

import (
	"regexp"
	"strings"
)

func FindIsogram(phrase string) bool {
	phrase = strings.ToLower(phrase)
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	for i := 0; i < len(phrase); i++ {
		if isAlpha(string(phrase[i])) {
			res := strings.Contains(phrase[i+1:], string(phrase[i]))
			if res {
				return false
			}
		}
	}
	return true
}
