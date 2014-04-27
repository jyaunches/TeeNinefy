package main

import(
	"testing"
)

func TestDictionaryFindsWords(t *testing.T) {
	dictionary := NewDictionary()
	mockWords := make(map[string][]string)
	mockWords["123"] = append(mockWords["123"], "dog") 
	mockWords["123"] = append(mockWords["123"], "cat") 

	dictionary.words = mockWords

	result := dictionary.lookUp("123")


	if (result[0] != "dog") || (result[1] != "cat") {
		t.Errorf("should return mocked words")
	}
}

func TestFindWordsReturnsEmptyArrayIfNothingFound(t *testing.T) {
	dictionary := NewDictionary()
	result := dictionary.lookUp("364")
	if len(result) != 0 {
		t.Errorf("should return empty array when no matches found")
	}
}

func TestProcessAndAddWordAddsToDictionarysWords(t *testing.T) {
	dictionary := NewDictionary()
	dictionary.processAndAddWord("dog")
	
	result := dictionary.lookUp("364")

	if (result[0] != "dog") {
		t.Errorf("should return mocked words")
	}
}