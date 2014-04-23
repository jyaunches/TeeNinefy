package main

import(
  "fmt"
  "os"
  "strconv"
)

func stringToInt(input string) int64 {
	intResult, err := strconv.ParseInt(input, 8, 64)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    return intResult
}