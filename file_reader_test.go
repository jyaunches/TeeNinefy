package main

import(
    "testing"
)

func TestFileReadingFromValidInputReturnsSameNumberEntriesAsFile(t *testing.T) {    
	letterMapper = NewLetterMapper()  	
	dictionary = make(map[string][]string)
    readFile("test_input.txt", ProcessLine)

    if(len(dictionary) != 3){
        t.Errorf("map size should be 3 was: ", len(dictionary))        
    }      
}

func TestFileReadingFromValidInputNumberRepMappedToWord(t *testing.T) {    
	letterMapper = NewLetterMapper()  	
	dictionary = make(map[string][]string)
    readFile("test_input.txt", ProcessLine)

    if(dictionary["27753"][0] != "apple"){
        t.Errorf("should have apple in mapping")        
    }      
}

func TestFileReadingShouldHandleWordsThatMapToSameNumber(t *testing.T) {    
	letterMapper = NewLetterMapper()  	
	dictionary = make(map[string][]string)
    readFile("test_input.txt", ProcessLine)

    if((dictionary["2337"][0] != "beer") || (dictionary["2337"][1] != "bees")){
        t.Errorf("should support words with same number mapping")        
    }      
}

// func TestFileReadingFromEmptyFile(t *testing.T) {    
//     wordArray := readFile("empty_input.txt")

//     if wordArray != nil {
//         t.Errorf("fail")        
//     }    
// }