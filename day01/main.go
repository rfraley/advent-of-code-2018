package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")
	changes := toIntSlice(lines)

	sum := finalFrequency(changes)
	fmt.Println("Sum:", sum)

	f := repeatedFrequency(changes)
	fmt.Println("First repeated frequency:", f)
}

func toIntSlice(ss []string) []int {
	si := []int{}
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		si = append(si, i)
	}
	return si
}

func finalFrequency(changes []int) int {
	sum := 0
	for _, i := range changes {
		sum += i
	}
	return sum
}

func repeatedFrequency(changes []int) int {
	m := make(map[int]bool)
	frequency := 0
	for {
		for _, change := range changes {
			frequency += change
			_, exists := m[frequency]
			if exists {
				return frequency
			}
			m[frequency] = true
		}
	}
}
