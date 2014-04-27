package main

import (
  "bufio"
  "os"
  "log"
)

func readFile(path string, processLine Processing)  {
  file, err := os.Open(path)
  if err != nil {
    log.Fatalf("Reading error: %s", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    processLine(scanner.Text())    
  }
}


