package main

import (
	"strings"	
)

type Dictionary struct {
	words []string
}

func (d *Dictionary) search(searchItems []string) []string {
	var results []string

	for j:= 0; j < len(searchItems); j++ {
		results = append(results, d.getWordsFor(searchItems[j])...)			
	}

	return results
}

func (d *Dictionary) getWordsFor(searchItem string) []string {
	var results []string
	for i := 0; i < len(d.words); i++ {		
		if strings.Index(d.words[i], searchItem) == 0 {
			results = append(results, d.words[i])			
		}
	}
	return results
}

func (d *Dictionary) addWords(newWords []string) {
	d.words = append(d.words, newWords...)
}