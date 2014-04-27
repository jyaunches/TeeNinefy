package main

import (
  "bufio"
  "os"
  "log"
  "path/filepath"
)

type FileReader struct {
    lineProcessor Processing
}

func NewFileReader(lP Processing) *FileReader {
    fileReader := &FileReader{}
    fileReader.lineProcessor = lP
    return fileReader
}

func (f *FileReader) processDirectory(path string) {
    filepath.Walk(path, f.visitDictionarySource) 
}

func (f *FileReader) visitDictionarySource(path string, info os.FileInfo, err error) error {
    file, err := os.Open(path)
    if err != nil {
      log.Fatalf("Reading error: %s", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      f.lineProcessor(scanner.Text())    
    }

    return err
}


