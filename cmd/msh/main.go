package main

import (
  "os"

  "github.com/justaskz/mashina/internal/cli"
)

func main() {
  cli.Init(os.Args[1:])
}
