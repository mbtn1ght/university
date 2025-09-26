package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numList := makeList()
	fmt.Print(sort(numList))

}

func makeList() []int {
	var s string
	var str string
	file, err := os.Open("numList.txt")
	if err != nil {
		fmt.Print("unavaluable object")
	}
	defer file.Close()
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)

		if n == 63 {
			break
		}

		if err != nil {
			break
		}

		str = string(data[:n])
		s += str

	}
	resultS := strings.Split(s, " ")

	var resaltN = []int{}

	for _, i := range resultS {
		j, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		resaltN = append(resaltN, j)
	}
	return resaltN

}

func sort(someList []int) []int {

	for {
		count := 0
		for i := 0; i < len(someList)-1; i++ {
			if someList[i] > someList[i+1] {
				someList[i], someList[i+1] = someList[i+1], someList[i]
				count += 1
			} else {
				continue
			}

		}
		if count == 0 {
			break
		}
	}
	return someList
}
