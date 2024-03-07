package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day1.txt")
	check(err)

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		check(err)

		numbers = append(numbers, num)
	}
	// Part 1
	// change between previous and next value
	var increase = 0

	for line, number := range numbers[1:] {
		if number > numbers[line] {
			increase++
		}
	}
	fmt.Println(increase)

	// Part 2
	var sum []int
	for line := 2; line < len(numbers); line++ {
		sum = append(sum, numbers[line]+numbers[line-1]+numbers[line-2])
	}

	increase = 0
	for line, number := range sum[1:] {
		if number > sum[line] {
			increase++
		}
	}
	fmt.Println(increase)
	file.Close()
}
