package main

import(
"testing"
// "fmt"
)

func TestNumberMapperMapsNumberToLetter(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.map_number("8")
    expected_characters := []string{"t", "u", "v"}
    if characters != nil {
    	for i := 0; i < len(characters); i++ {	
    		if characters[i] != expected_characters[i] {
				t.Errorf("fail")
    		}
    	}
    }
}
