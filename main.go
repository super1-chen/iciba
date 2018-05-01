package main

import (
	"flag"
	"fmt"
	"os"
	"iciba/searcher"
)

func main() {

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Printf("Usage: %s 'words'", os.Args[0])
		os.Exit(1)
	}
	words := args[0]
	err := searcher.SerchWords(words)
	if err != nil {
		fmt.Println("啊哦，好像出错了，try again~")
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
