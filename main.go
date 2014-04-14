package main

import (
  "fmt"
  "os"
)

func main() {
  lines := readFile("american-words.80")
  dictionary := NewDictionary(lines)
  mapper := NewNumberMapper()
  
  input := os.Args[1]
  characters := mapper.mapNumber(input)  
  results := dictionary.search(characters)   	
  fmt.Println(results) 
}



