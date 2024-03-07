package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct{
	comm string
	attr int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./day2.txt")
	check(err)

	var commands []command
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		comm := parts[0]
		attr,err := strconv.Atoi(parts[1])
		check(err)
		
		commands = append(commands,
			command{
				comm: comm,
				attr: attr,
			})
	}
	

	// part 1
	var horizontal, depth int
	for _, cmd := range commands{
		switch cmd.comm{
		case "up":
			depth-=cmd.attr
		case "down":
			depth+=cmd.attr
		case "forward":
			horizontal+=cmd.attr
		}

	}
	fmt.Println(horizontal * depth)

	// part 2
	horizontal = 0
	depth = 0
	var aim int
	for _, cmd := range commands{
		switch cmd.comm{
		case "up":
			aim-=cmd.attr
		case "down":
			aim+=cmd.attr
		case "forward":
			horizontal+=cmd.attr
			depth += aim * cmd.attr
		}
	}
	fmt.Println(horizontal * depth)

	file.Close()
}
