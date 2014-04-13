package main

import "testing"

func TestDictionaryFindsWordsBasedOnFirstLetter(t *testing.T) {
    mock_words := []string{"apple", "pear"}
    d := NewDictionary(mock_words)
    results := d.search([]string{"p", "q"})
    if results == nil || results[0] != "pear"{
    	t.Errorf("fail")    	
    }
}

func TestDictionaryFindsWordsBasedOnFirstTwoLetters(t *testing.T) {
    mock_words := []string{"gone", "after", "pear"}
    d := NewDictionary(mock_words)
    results := d.search([]string{"af", "ad"})
    if len(results) != 1{
    	t.Errorf("should have 1 result")    	
    }

    if results[0] != "after" {
    	t.Errorf("result not as expected")    	
    }
}

func TestDictionaryFindsConsecutiveWordsForSameSearchString(t *testing.T) {
    mock_words := []string{"gone", "after", "afore", "pear"}
    d := NewDictionary(mock_words)
    results := d.search([]string{"af", "ad"})
    if len(results) != 2{
    	t.Errorf("should have 2 results")    	
    }

    if results[0] != "after" && results[0] != "afore"{
    	t.Errorf("should return consecutive results for same search string")    	
    }
}

func TestDictionaryReturnsEmptyResultsIfNoWordsFound(t *testing.T) {
    mock_words := []string{"apple", "pear"}
    d := NewDictionary(mock_words)
    results := d.search([]string{"d"})

    if len(results) != 0 {
    	t.Errorf("fail")    
    }    
}