// run this with `go run` (golang must be installed)
package main

import "fmt"

func main() {
  flag := "c8e49=<@"
  arr := []rune(flag)

  fmt.Printf("flag: %s\n", flag);
  fmt.Printf("decoded: ")
  for i, el := range arr {
    fmt.Printf("%c", el - rune(i))
  }
  fmt.Printf("\n")
}
