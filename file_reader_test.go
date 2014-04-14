package main

import(
    "testing"
)

func TestFileReadingFromValidInput(t *testing.T) {    
    wordArray := readFile("test_input.txt")

    if wordArray == nil || wordArray[0] != "apple" || wordArray[1] != "orange" {
        t.Errorf("fail")        
    }    
}

func TestFileReadingFromEmptyFile(t *testing.T) {    
    wordArray := readFile("empty_input.txt")

    if wordArray != nil {
        t.Errorf("fail")        
    }    
}