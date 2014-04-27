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
	fmt.Println("Enter 'Done' when you are finished!")
	fmt.Println("")	
}

func getUserInput() string{
  fmt.Printf("Enter your numbers:")
  var input string	
	_, err := fmt.Scanf("%s", &input)
  if err !=nil {
    log.Fatalf("Reading error: %s", err)   
  }
  return input
}

func printResult(result string){
	fmt.Printf("Results: ")
  	fmt.Printf(result)
	fmt.Println("\n")  
}

func printGoodbye(){
	fmt.Println("Thanks! Goodbye.")	
}