package main

import (
	"regexp"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// current processs
type StringChecker func(string) FreqMap

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.

func Frequency(s string) FreqMap {
	m := FreqMap{}
	isAlpha := regexp.MustCompile(`^[A-Za-z0-9]*$`).MatchString
	for _, r := range s {
		if isAlpha(string(r)) {
			m[r]++
		}
	}
	return m
}

func ConcurrentFrequency(l []string) FreqMap {
	res := FreqMap{}
	ch := make(chan FreqMap, 10)

	var wg sync.WaitGroup

	wg.Add(len(l))
	for _, r := range l {
		go func(t string) {
			ch <- Frequency(t)
			wg.Done()
		}(r)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for FreqMap := range ch {
		for letter, freq := range FreqMap {
			res[letter] += freq
		}
	}
	return res
}
