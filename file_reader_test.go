package main

import(
    "testing"
)

func TestFileReadingFromValidInput(t *testing.T) {    
    word_array := readFile("test_input.txt")

    if word_array == nil || word_array[0] != "apple" || word_array[1] != "orange" {
        t.Errorf("fail")        
    }    
}

func TestFileReadingFromEmptyFile(t *testing.T) {    
    word_array := readFile("empty_input.txt")

    if word_array != nil {
        t.Errorf("fail")        
    }    
}