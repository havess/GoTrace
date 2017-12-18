package system

import "fmt"

func Error(s string) {
    fmt.Printf("!!! Anvil has encountered an error: %s !!!\n", s)
}
