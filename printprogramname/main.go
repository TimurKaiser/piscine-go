package main

import (
	"os.Args"
)

func main() {
	args := os.Args
	if len(args) > 0 {
		return args
	}
}
