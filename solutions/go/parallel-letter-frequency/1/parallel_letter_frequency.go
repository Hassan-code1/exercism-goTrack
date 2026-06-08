package parallelletterfrequency

import (
	"unicode"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	mp := make(FreqMap)
	for _, ch := range text {
		if(unicode.IsLetter(ch)){
			ch := unicode.ToLower(ch)
			mp[rune(ch)]++
		}
	}
	return mp
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	ch := make(chan FreqMap, len(texts))
	for _, text := range texts {
		go func(text string){
			ch <- Frequency(text)
		}(text)
	}
	result := make(FreqMap) 
	for i := 0; i < len(texts); i++ {
    	fm := <-ch

    	for r, count := range fm {
       	 result[r] += count
    	}
	}
	return result
}
