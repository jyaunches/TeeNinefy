package main

import(
)

type NumberMapper struct {
	number_map map[rune][]string
}

func NewNumberMapper() *NumberMapper {	
	return &NumberMapper{number_map : map[rune][]string{
			'0': {},
		    '1': {},
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

func (d *NumberMapper) map_number(numbers_entered string) []string {	
	var letter_combinations []string

    for _, digit := range numbers_entered{
    	if(len(letter_combinations) == 0){
    		letter_combinations = d.build_combinations("", digit)
		}else{
			var next_combinations []string
			for i := 0; i < len(letter_combinations); i++ {	
				next_combinations = append(next_combinations, d.build_combinations(letter_combinations[i], digit)...)
			}
			letter_combinations = next_combinations			
		}
    }
    return letter_combinations
}

func (d *NumberMapper) build_combinations(base string, next_number rune) []string {	
	next_letters := d.number_map[next_number]

	results := make([]string, len(next_letters)) 
	for i := 0; i < len(next_letters); i++ {	
    	results[i] = base + next_letters[i]
    }
	return results
}