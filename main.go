package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}

		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			fmt.Println("Hexdump erlang lists/strings, e.g. [72,101,108,108,111]")
			fmt.Println("The command works with pipes or files.")
			fmt.Println("Usage: cat file | erldump")
			fmt.Println("Usage: erldump file")
			return
		}

		parse(bufio.NewReader(os.Stdin))
		return
	}

	for _, fn := range os.Args[1:] {
		if reader, err := os.Open(fn); err == nil {
			parse(reader)
		}
	}
}

func parse(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	src := make([]byte, 0)
	rex := regexp.MustCompile(`[^0-9]+`)

	for scanner.Scan() {
		for _, n := range rex.Split(scanner.Text(), -1) {
			if i, err := strconv.Atoi(n); err == nil {
				src = append(src, uint8(i))
			}
		}
	}

	fmt.Println(hex.Dump(src))
}
