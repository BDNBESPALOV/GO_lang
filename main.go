package main

import (
    "fmt"
     "os"
     )

func main() {
    s, str :="", ""
	for _, arg := range os.Args[1:] {
	s += str + arg
	    fmt.Print(s)
	}
}
