package main

import "os"

func main() {
	os.Stderr.WriteString("hello world")
}
