package main

import (
	"errors"
)

func GetHammingDistance(firstSeq, secondSeq string) (int, error) {
	if len(firstSeq) != len(secondSeq) {
		return int(0), errors.New("strings are not equal length")
	} else {
		count := 0
		for i := 0; i < len(firstSeq); i++ {
			if firstSeq[i] != secondSeq[i] {
				count++
			}
		}
		return count, nil
	}
}
