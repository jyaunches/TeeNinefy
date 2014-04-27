package main

import(
)

type Dictionary struct {
	words map[string][]string
	letterMapper *LetterMapper
}

func NewDictionary() *Dictionary {	
	dictionary := &Dictionary{}
	dictionary.words = make(map[string][]string)
	dictionary.letterMapper = NewLetterMapper()		
	return dictionary
}

func (d *Dictionary) lookUp(input string) string{
	return d.words[input][0]
}

type Processing func(string) 

func (d *Dictionary) processAndAddWord(line string) {
  lineAsNumber := d.letterMapper.mapWord(line)
  d.words[lineAsNumber] = append(d.words[lineAsNumber], line)        
}