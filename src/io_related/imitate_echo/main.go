package main

import (
	"flag" // command line option parser
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	// flag.Arg(0) 就是第一个真实的 flag
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
