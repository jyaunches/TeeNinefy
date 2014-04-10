package main

import (
  "fmt"
)

type CSVParser func(fileName string) []byte

func main() {
	lines := readFile("american-words.80")


 	 
	for i, line := range lines {
    	fmt.Println(i, line)
  }
}



