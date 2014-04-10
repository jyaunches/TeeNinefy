package main

import(
    "testing"
)

func TestFileReading(t *testing.T) {    
    word_array := readFile("test_input.txt")

    if word_array == nil || word_array[0] != "apple" || word_array[1] != "orange" {
        t.Errorf("fail")        
    }    
}
