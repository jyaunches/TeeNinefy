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

func (d *Dictionary) search(search_items []string) []string {
	var results []string

	for j:= 0; j < len(search_items); j++ {
		results = append(results, d.get_words_for(search_items[j])...)			
	}

	return results
}

func (d *Dictionary) get_words_for(search_item string) []string {
	var results []string
	for i := 0; i < len(d.words); i++ {		
		if strings.Index(d.words[i], search_item) == 0 {
			results = append(results, d.words[i])			
		}
	}
	return results
}