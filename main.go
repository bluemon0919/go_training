package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n bool // 行番号有無
func init() {
	flag.BoolVar(&n, "n", false, "行番号有無（デフォルト無し）")
}
func main() {
	ln := 0
	flag.Parse()
	fs := flag.Args()
	for _, f := range fs {
		fileOutput(f, &ln, n)
	}
}

func fileOutput(filePath string, line *int, b bool) {
	sf, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer sf.Close()
	scanner := bufio.NewScanner(sf)
	for scanner.Scan() {
		var s string
		if b {
			*line++
			s = fmt.Sprintf("%d: ", *line)
		}
		fmt.Printf("%s%s\n", s, scanner.Text())
	}
}
