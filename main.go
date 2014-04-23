package main

import (
  "fmt"
  "os"
  "path/filepath"
)

var input string
var dictionary map[string][]string
var letterMapper *LetterMapper

type Processing func(string) 

func ProcessLine(line string) {
  lineAsNumber := letterMapper.mapWord(line)
  dictionary[lineAsNumber] = append(dictionary[lineAsNumber], line)        
}

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("Please provide your input! Ex. 4663 -> home\n")
    return
  }

  letterMapper = NewLetterMapper()  
  dictionary = make(map[string][]string)

  input = os.Args[1]
  filepath.Walk("dictionary_sources", visitDictionarySource)    
  fmt.Printf("Result: ")
  fmt.Printf(dictionary[input][0])
  fmt.Printf("\n")
}

func visitDictionarySource(path string, f os.FileInfo, err error) error {
  readFile(path, ProcessLine)  
  return err
}



