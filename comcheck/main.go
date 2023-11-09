package main

import "os"

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			os.Stdout.WriteString("Alert!!!\n")
			break
		}
	}
}
