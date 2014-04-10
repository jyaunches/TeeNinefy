package main

import (
  "bufio"
  "os"
  "log"
)

func readFile(path string) []string {
  file, err := os.Open(path)
  if err != nil {
    log.Fatalf("Reading error: %s", err)
    return nil
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines
}

