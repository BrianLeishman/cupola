package main

import "flag"

func main() {
	var input string
	flag.StringVar(&input, "i", "", "Set an input PST file location")
	flag.Parse()

}
