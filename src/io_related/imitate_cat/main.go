package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
	}
	return
}
func main() {
	flag.Parse()
	// 没有args 就从终端读
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	// 有args，就从文件中读
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}
