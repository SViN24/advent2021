package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day3.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	var one_counter [12]int
	var line_count int

	for scanner.Scan() {
		bLine := scanner.Text()
		line_count++

		for i := range bLine {
			if bLine[i] == '1' {
				one_counter[i]++
			}
		}
	}

	// part 1

	var gamma_spectrum, epsilon_spectrum int

	for i := 0; i < 12; i++ {
		if one_counter[i] > line_count-one_counter[i] {
			gamma_spectrum |= 1 << (11 - i)
		} else {
			epsilon_spectrum |= 1 << (11 - i)
		}
	}

	fmt.Println(gamma_spectrum * epsilon_spectrum)

	file.Close()
}
