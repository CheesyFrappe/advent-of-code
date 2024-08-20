package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var numberWords = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	filename := "input.txt"

	file, _ := os.Open(filename)
	defer file.Close()

	part_one(file)
	file.Seek(0, 0)
	part_two(file)
}

func part_one(file *os.File) {
	r := bufio.NewReader(file)
	total := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		var arr []int

		for _, char := range line {
			if unicode.IsDigit(rune(char)) {
				arr = append(arr, int(char-'0'))
			}
		}
		first_val := arr[0]
		second_val := arr[len(arr)-1]

		total += (first_val * 10) + second_val
	}

	fmt.Println(total)
}

func part_two(file *os.File) {
	defer file.Close()

	r := bufio.NewReader(file)
	total := 0

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		var arr []int
		for i := 0; i < len(line); i++ {
			for k, v := range numberWords {
				if string(line[i:i+len(k)]) == k {
					arr = append(arr, v)
					break
				}
			}
			if unicode.IsDigit(rune(line[i])) {
				arr = append(arr, int(line[i]-'0'))
			}
		}

		first_val := arr[0]
		second_val := arr[len(arr)-1]

		total += (first_val * 10) + second_val
	}
	fmt.Println(total)
}
