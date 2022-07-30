package main

import (
	"os"
	"testing"
)

var (
	texts  = []string{"text_1.txt", "text_2.txt", "text_3.txt", "text_4.txt", "text_5.txt"}
	chunks []string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSequenceTexts(texts ...string) string {
	var seq string
	for _, text := range texts {
		chunk, err := os.ReadFile(text)
		check(err)
		seq += string(chunk)
	}
	return seq
}
func getSliceTexts(texts ...string) []string {
	var routines []string
	for _, text := range texts {
		chunk, err := os.ReadFile(text)
		check(err)
		chunks = append(chunks, string(chunk))
	}
	return routines
}
func BenchmarkFrequency(b *testing.B) {
	total_texts := getSequenceTexts(texts...)
	//getSequenceTexts(texts...)

	Frequency(total_texts)
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	routines := getSliceTexts(texts...)

	ConcurrentFrequency(routines)
}
