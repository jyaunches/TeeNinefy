package main

import "testing"

func TestDictionary(t *testing.T) {
    mock_words := []string{"apple", "pear"}
    d := NewDictionary(mock_words)
    d.search_number("490")
    t.Errorf("fail")
}
