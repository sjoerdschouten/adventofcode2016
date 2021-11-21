package utils

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads all lines from a provided filepath and return a slice of strings
// if the filepath does not exist it will return an error.
func ReadLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, fmt.Errorf("ReadLines(): %s", err)   
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var lines []string
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, nil
}

func ReadEntireFile(path string) string {
  result, err := os.ReadFile(path)
  if err != nil {
    panic(err)
  }
  return string(result)
}
