package parser

import (
    "fmt"
)

func ParseFile(s string) bool {
    fmt.Printf("Starting Parser...\n")

    if s == "-" {
        fmt.Printf("Parsing from standard input")
    }
    
    return true
}
