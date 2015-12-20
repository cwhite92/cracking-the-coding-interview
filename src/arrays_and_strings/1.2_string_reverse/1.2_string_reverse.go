package main

import (
  "fmt"
)

func main() {
  strings := []string {"Hello, world!", "hannah", "à la carte"}

  for _, string := range strings {
    fmt.Printf("\"%s\" reversed is: \"%s\"\n", string, reverse(string))
  }
}

func reverse(input string) string {
  runes := []rune(input)

  // i and j start at oposite ends of the string, and on each iteration they
  // will swap runes and work their way towards the middle of the string.
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }

  return string(runes)
}
