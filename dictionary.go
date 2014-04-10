package main

import (	
	"strings"	
)

type Dictionary struct {
	words []string
}

func NewDictionary(source_words []string) *Dictionary {	
	return &Dictionary{words : source_words}
}

func (d *Dictionary) search_number(search_string string) []string {
	var results []string
	for i := 0; i < len(d.words); i++ {		
		if strings.Index(d.words[i], search_string) == 0 {
			results = append(results, d.words[i])
		}
	}

	return results
}