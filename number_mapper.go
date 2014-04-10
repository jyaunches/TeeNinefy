package main

type NumberMapper struct {
	number_map map[string][]string
}

func NewNumberMapper() *NumberMapper {	
	return &NumberMapper{number_map : map[string][]string{
			"0": {},
		    "1": {},
		    "2": {"a", "b", "c"},
		    "3": {"d", "e", "f"},
		    "4": {"g", "h", "i"},
		    "5": {"j", "k", "l"},
		    "6": {"m", "n", "o"},
		    "7": {"p", "q", "r", "s"},
		    "8": {"t", "u", "v"},
		    "9": {"w", "x", "y", "z"},
	}}
}

func (d *NumberMapper) map_number(single_number string) []string {	
	return d.number_map[single_number]
}