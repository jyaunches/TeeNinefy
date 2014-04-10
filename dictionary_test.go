package main

import "testing"

func TestDictionaryFindsWordsBasedOnFirstLetter(t *testing.T) {
    mock_words := []string{"apple", "pear"}
    d := NewDictionary(mock_words)
    results := d.search_number("p")
    if results == nil || results[0] != "pear"{
    	t.Errorf("fail")    	
    }
}

func TestDictionaryReturnsEmptyResultsIfNoWordsFound(t *testing.T) {
    mock_words := []string{"apple", "pear"}
    d := NewDictionary(mock_words)
    results := d.search_number("d")

    if len(results) != 0 {
    	t.Errorf("fail")    
    }    
}