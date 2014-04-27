package main

import (
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
  letterMapper = NewLetterMapper()  
  dictionary = make(map[string][]string)
  filepath.Walk("dictionary_sources", visitDictionarySource)    

  printGreeting()

  input := getUserInput()

  for ; ((input != "Done") && (input != "done")); {    
    printResult(dictionary[input][0])
    input = getUserInput()
  }  
  printGoodbye()  
}

func visitDictionarySource(path string, f os.FileInfo, err error) error {
  readFile(path, ProcessLine)  
  return err
}



