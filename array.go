package main

import (
    "fmt"
    "sort"
)

func main () {
  a := []string{
    "shit",
    "fuck",
    "ass",
    "dick",
    "cunt",
    "whore",
    "fag",
    "dammit",
    "crack",
    "pipe",
  }

  // Print the Array
  fmt.Println(a)
  // Print the length of the Array
  fmt.Println(len(a))

  // Sort the Array then print
  sort.Strings(a)
  fmt.Println(a)

}
