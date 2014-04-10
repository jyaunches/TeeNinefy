package main

type Dictionary struct {
	word []string
}

func NewDictionary(source_words []string) *Dictionary {
	return new(Dictionary)	
}

func (d *Dictionary) search_number(search_string string) []string{
	return []string {}
}