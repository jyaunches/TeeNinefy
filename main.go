package main

import (
  "fmt"
  "os"
)

type CSVParser func(fileName string) []byte

func main() {
  lines := readFile("american-words.80")
  dictionary := NewDictionary(lines)
  mapper := NewNumberMapper()
  
  input := os.Args[1]
  characters := mapper.map_number(input)  
  results := dictionary.search(characters)   	
  fmt.Println(results) 
}



