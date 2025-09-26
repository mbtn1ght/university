package main

import (
	"fmt"
	"os"
)

func main() {
	in, err := os.Open("INPUT.TXT")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var k int
	var s string
	_, err = fmt.Fscan(in, &k, &s)
	if err != nil {
		panic(err)
	}

	// k начинается с 1, а индексы в Go с 0
	result := s[:k-1] + s[k:]

	out, err := os.Create("OUTPUT.TXT")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	fmt.Fprint(out, result)
}
