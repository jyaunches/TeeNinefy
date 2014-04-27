package main

func main() {    
  printGreeting()

  dictionary := NewDictionary()
  fileReader := NewFileReader(dictionary.processAndAddWord)
  fileReader.processDirectory("dictionary_sources")

  input := getUserInput()

  for ; ((input != "Done") && (input != "done")); {    
    printResult(dictionary.lookUp(input))
    input = getUserInput()
  }  
  printGoodbye()  
}





