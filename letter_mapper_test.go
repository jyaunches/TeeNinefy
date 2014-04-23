package main

import(
"testing"
)

func TestLetterMapperReturnsSameSizeAsInput(t *testing.T) {
    mapper := NewLetterMapper()
    wordAsNumbers := mapper.mapWord("fun")
	if len(wordAsNumbers) != 3 {
		t.Errorf("Wrong length of result: %d", len(wordAsNumbers))
    }
}

func TestLetterMapperReturnsExpectedStringOfNumbers(t *testing.T) {
    mapper := NewLetterMapper()
    wordAsNumbers := mapper.mapWord("fun")
	if wordAsNumbers != "486" {
		t.Errorf("Wrong result")
    }
}