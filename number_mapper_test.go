package main

import(
"testing"
)

func TestNumberMapperReturnsAnArray(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.mapNumber("8")
	if characters == nil {
		t.Errorf("array not returned")
    }
}

func TestNumberMapperReturnsCorrectNumberOfResults(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.mapNumber("8")
    
    if len(characters) != 3{
    	t.Errorf("Expected 3 results, got %d", len(characters))
    }
}

func TestNumberMapperReturnsCorrectResultsForNumber(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.mapNumber("8")
    expectedCharacters := []string{"t", "u", "v"}

	for i := 0; i < len(characters); i++ {	
		if characters[i] != expectedCharacters[i] {
			t.Errorf("fail")
		}
	}
}    	

func TestNumberReturnsCombinationsForStringAndArrayOfRunes(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.buildCombinations("to", '2')
    expectedCharacters := []string{"toa", "tob", "toc"}
	if len(characters) != 3{
    	t.Errorf("Expected 3 results, got %d", len(characters))
    }

    for i := 0; i < len(characters); i++ {	    	
		if characters[i] != expectedCharacters[i] {
			t.Errorf("incorrect combinations returned")
		}
	}
}

func TestNumberMapperMapsTwoNumbersToLetterPossibilities(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.mapNumber("82")
    expectedCharacters := []string{"ta", "tb", "tc", "ua", "ub", "uc", "va", "vb", "vc"}
    if len(characters) != 9{
    	t.Errorf("Expected 9 results, got %d", len(characters))
    }
    
	for i := 0; i < len(characters); i++ {	    	
		if characters[i] != expectedCharacters[i] {
			t.Errorf("incorrect combinations returned")
		}
	}
}

func TestNumberMapperMapsThreeNumbersToLetterPossibilities(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.mapNumber("827")
    expectedCharacters := []string{"tap","taq","tar","tas", 
    								"tbp","tbq","tbr","tbs", 
    								"tcp","tcq","tcr","tcs", 
    								"uap","uaq","uar","uas", 
    								"ubp","ubq","ubr","ubs", 
    								"ucp","ucq","ucr","ucs", 
    								"vap","vaq","var","vas", 
    								"vbp","vbq","vbr","vbs", 
    								"vcp","vcq","vcr","vcs"}
    if len(characters) != 36{
    	t.Errorf("Expected 36 results, got %d", len(characters))
    }
    
	for i := 0; i < len(characters); i++ {	    	
		if characters[i] != expectedCharacters[i] {
			t.Errorf("incorrect combinations returned")
		}
	}
}

func TestNumberMapperStripsInputOfOnesAndZeros(t *testing.T) {
    mapper := NewNumberMapper()
    characters := mapper.mapNumber("71400")

    if len(characters) != 12{
    	t.Errorf("Should only have combinations for 7 and 4.. got %d results", len(characters))
    }
}

