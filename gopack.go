package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("GoPack 0.0: Useless Package Manager at this point\n")
	fmt.Printf("Usage: %s [inputfile]\n", os.Args[0])
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	for i := 1; i < len(args)+1; i++ {
		fmt.Printf("param: %s\n", os.Args[i])
	}
}
