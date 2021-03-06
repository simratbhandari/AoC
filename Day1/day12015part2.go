package main

import (
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./day-01-input.txt")
	if err != nil {
		panic(err)
	}

	string := string(input)

	var floor int

	for index, char := range string {
		if string(char) == string("(") {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			println(index + 1)
			return
		}
	}
}