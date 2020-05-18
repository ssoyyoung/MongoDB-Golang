package main

import (
"os"
"fmt"
)

func main() {
    path, _ := os.Getwd()
    fmt.Println(path)
}
