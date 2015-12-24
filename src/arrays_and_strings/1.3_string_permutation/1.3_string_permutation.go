package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  var words = []struct {
    str1 string
    str2 string
  }{
    {"permutation", "unoptmetira"},
    {"hello world", "oellh owlrd"},
    {"ぁあぃぎ", "じぎぁあぃ"},
    {"", ""},
    {"()[]", "][)(!"},
  }

  for _, c := range words {
    fmt.Printf("Is \"%s\" a permutation of \"%s\"? %t\n", c.str2, c.str2, permutation(c.str1, c.str2))
  }
}

func permutation(str1 string, str2 string) bool {
  if utf8.RuneCountInString(str1) != utf8.RuneCountInString(str2) {
    // Strings are not the same length - quick win
    return false
  }

  // Add each character of str1 to a map
  // key: the character
  // value: the number of occurances of that character
  charMap := make(map[rune]int)

  for _, c := range str1 {
    charMap[c]++
  }

  // Decrement each character of charMap when that character is seen in str2
  for _, c := range str2 {
    if _, ok := charMap[c]; ok {
      charMap[c]--
    } else {
      // This character does not appear in str1
      return false
    }
  }

  // Check if map has any value != 0
  for _, count := range charMap {
    if (count != 0) {
      return false;
    }
  }

  return true
}
