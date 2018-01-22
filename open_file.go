package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  lines := openFile("stat.log")
  fmt.Printf("lines: %v\n", lines)
}

func openFile(path string) []string {
  f, err := os.Open(path)
  if err != nil {
    fmt.Fprintf(os.Stderr, "open error %v\n", err)
    os.Exit(1)
  }

  defer f.Close()

  lines := []string{}
  scan := bufio.NewScanner(f)
  for scan.Scan() {
    lines = append(lines, scan.Text())
  }
  if serr := scan.Err(); serr != nil {
    fmt.Fprintf(os.Stderr, "scan error %v\n", serr)
  }

  return lines
}
