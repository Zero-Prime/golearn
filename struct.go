package main

import (
  "fmt"
  "sort"
)

type Foo struct {
  Value  string
}

type Foos []Foo

func (slice Foos) Len() int {
	return len(slice)
}

func (slice Foos) Less(i, j int) bool {
	return slice[i].Value > slice[j].Value;
}

func (slice Foos) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main () {
  foos := Foos{
    {Value: "shit"},
    {Value: "fuck"},
    {Value: "ass"},
    {Value: "dick"},
    {Value: "cunt"},
    {Value: "whore"},
    {Value: "fag"},
    {Value: "dammit"},
    {Value: "crack"},
    {Value: "pipe"},
  }

  sort.Sort(foos)

  for i, c := range foos {
    fmt.Println(i, c.Value)
  }
}
