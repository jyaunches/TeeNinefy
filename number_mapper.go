// The responsibility of NumberMapper is:
// #1 - to hold the number to character values
// #2 - given a string of numbers to map to all the potential 
// character combinations a word can start with and return that 
// as an array

package main

import(
	"strings"
)

type NumberMapper struct {
	numberMap map[rune][]string
}

func NewNumberMapper() *NumberMapper {	
	return &NumberMapper{numberMap : map[rune][]string{
		    '2': {"a", "b", "c"},
		    '3': {"d", "e", "f"},
		    '4': {"g", "h", "i"},
		    '5': {"j", "k", "l"},
		    '6': {"m", "n", "o"},
		    '7': {"p", "q", "r", "s"},
		    '8': {"t", "u", "v"},
		    '9': {"w", "x", "y", "z"},
	}}
}

func (d *NumberMapper) mapNumber(numbersEntered string) []string {	
	r := strings.NewReplacer("0", "", "1", "")
	numbersEntered = r.Replace(numbersEntered)

	var letterCombinations []string

    for _, digit := range numbersEntered{
    	if(len(letterCombinations) == 0){
    		letterCombinations = d.buildCombinations("", digit)
		}else{
			var nextCombinations []string
			for i := 0; i < len(letterCombinations); i++ {	
				nextCombinations = append(nextCombinations, d.buildCombinations(letterCombinations[i], digit)...)
			}
			letterCombinations = nextCombinations			
		}
    }
    return letterCombinations
}

func (d *NumberMapper) buildCombinations(base string, nextNumber rune) []string {	
	nextLetters := d.numberMap[nextNumber]

	results := make([]string, len(nextLetters)) 
	for i := 0; i < len(nextLetters); i++ {	
    	results[i] = base + nextLetters[i]
    }
	return results
}