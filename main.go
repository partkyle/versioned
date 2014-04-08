package main

import (
	"flag"
	"fmt"
)

var version string

var (
	printVersion = flag.Bool("version", false, "print version and exit")
)

func main() {
	flag.Parse()

	if *printVersion {
		fmt.Printf("version: %s\n", version)
		return
	}

	fmt.Printf("running application...\n")
}
