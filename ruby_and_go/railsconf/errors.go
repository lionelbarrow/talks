package main

import (
  "fmt"
  "math/rand"
  "errors"
)

func ComplicatedOperation(input string) (string, error) {
  if rand.Rand() > 0.5 {
    return "Success!", nil
  }
  return "", errors.New("Not enough bananas")
}

func main() {
  output, err := ComplicatedOperation("input")
  if err != nil {
    fmt.Println("There was an error! It was: " + err.Error())
  } else {
    fmt.Println(output)
  }
}
