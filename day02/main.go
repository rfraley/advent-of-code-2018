package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\r\n")

	checksum := calcChecksum(lines)
	fmt.Println("Checksum:", checksum)

	common := findIt(lines)
	fmt.Println("Same:", common)
}

func calcChecksum(ids []string) int {
	twoCountSum := 0
	threeCountSum := 0

	for _, id := range ids {
		twoFound, threeFound := includeID(id)
		if twoFound {
			twoCountSum++
		}
		if threeFound {
			threeCountSum++
		}
	}
	return twoCountSum * threeCountSum
}

func includeID(id string) (bool, bool) {
	m := make(map[rune]int)
	for _, c := range id {
		count, _ := m[c]
		m[c] = count + 1
	}

	hasTwoCount := false
	hasThreeCount := false
	for _, v := range m {
		hasTwoCount = hasTwoCount || v == 2
		hasThreeCount = hasThreeCount || v == 3
		if hasTwoCount && hasThreeCount {
			break
		}
	}
	return hasTwoCount, hasThreeCount
}

func findIt(ids []string) string {
	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			same := sameLetters(ids[i], ids[j])
			if (len(ids[i]) - len(same)) == 1 {
				return same
			}
		}
	}
	panic("not found")
}

func sameLetters(a string, b string) string {
	var sb strings.Builder
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			sb.WriteByte(a[i])
		}
	}
	return sb.String()
}
