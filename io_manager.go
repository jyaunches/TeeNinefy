package main

import (
  "fmt"
  "log"
)

func printGreeting() {
	fmt.Println("")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Welcome to the amazing T9 simulator program.") 
	fmt.Println("Enter any series of numbers and I will find you word matches.")
	fmt.Println("")	
}

func getUserInput() string{
  fmt.Printf("Type your numbers (or 'done'):")
  var input string	
	_, err := fmt.Scanf("%s", &input)
  if err !=nil {
    log.Fatalf("Reading error: %s", err)   
  }
  return input
}

func printResult(results []string){
	if len(results) == 0{
		fmt.Printf("No results found!")
	}else{
		fmt.Printf("Results: ")
		for i := range results {
        	fmt.Printf("%s ", results[i])
		}				
	}  	
	fmt.Println("\n")  
}

func printGoodbye(){
	fmt.Println("Thanks! Goodbye.")	
}