package main

import "strings"

func FindIsogram(phrase string) (bool, error) {
	var res bool
	for i := 0; i < len(phrase); i++ {
		res = strings.Contains(phrase, string(phrase[i]))
		print(string(phrase[i]))
		if res {
			return false, nil
		}
	}
	return true, nil
}
