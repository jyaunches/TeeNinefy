
package main

type LetterMapper struct {
	letterMap map[rune]string
}

func NewLetterMapper() *LetterMapper {	
	return &LetterMapper{letterMap : map[rune]string{
		'a':"2", 'b':"2", 'c':"2",
		'd':"3", 'e':"3", 'f':"4",
		'g':"4", 'h':"4", 'i':"4",
		'j':"5", 'k':"5", 'l':"5",
		'm':"6", 'n':"6", 'o':"6",
		'p':"7", 'q':"7", 'r':"7", 's':"7",		
		't':"8", 'u':"8", 'v':"8",
		'w':"9", 'x':"9", 'y':"9", 'z':"9",				
	}}
}

func (d *LetterMapper) mapWord(word string) string {
	var result string
	for _, letter := range word{
		if val, ok := d.letterMap[letter]; ok {
			result = result + val
	    }
	}
	
    return result
}