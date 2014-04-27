package main

import(
    "testing"
)

var linesAdded []string

func processWordMock(line string) {
	linesAdded = append(linesAdded, line)	  
}

func TestDirectoryReadingProcessesUsingProcessor(t *testing.T) {
	fileReader := NewFileReader(processWordMock)

	if len(linesAdded) != 0{
        t.Errorf("linesAdded should be 0 before starting")        
	}

	fileReader.processDirectory("test_input")
	
	if len(linesAdded) != 5{
        t.Errorf("linesAdded should be 5 after processing")        
	}
}

