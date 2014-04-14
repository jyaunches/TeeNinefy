package main

import (
  "fmt"
  "os"
  "path/filepath"
)

var dictionary Dictionary

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Please provide your input! Ex. 4663 -> home\n")
    return
  }

  filepath.Walk("dictionary_sources", visitDictionarySource)    
  mapper := NewNumberMapper()
     
  characters := mapper.mapNumber(os.Args[1])  
  results := dictionary.search(characters)
  fmt.Println("Results:")    	
  fmt.Println(results) 
}

func visitDictionarySource(path string, f os.FileInfo, err error) error {
  lines := readFile(path)
  dictionary.addWords(lines)
  return err
}



