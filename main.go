package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	name := os.Args[1]
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	index := 0
	var ascii string
	for {

		if index%8 == 0 && index != 0 {
			fmt.Printf(" ")
		}
		if index%16 == 0 {
			if index > 0 {
				fmt.Printf("|%s|", ascii)
				ascii = ""
				fmt.Printf("\n")
			}
			fmt.Printf("%08x  ", index)
		}
		b, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			if index%16 != 0 {
				fmt.Printf("|%s|\n", ascii)
				fmt.Printf("%08x  ", index)
			}
			fmt.Printf("\n")
			break
		}

		index++

		char := string(b)
		if b < 0x1F || b > unicode.MaxASCII {
			char = "."
		}
		ascii = fmt.Sprintf("%s%s", ascii, char)

		fmt.Printf("%02x ", b)
	}
}
