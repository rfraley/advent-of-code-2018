package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type claim struct {
	id       int
	start    position
	end      position
	size     size
	overlaps bool
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(data), "\r\n")
	claims := []*claim{}
	points := make(map[string][]*claim)
	totalOverlap := 0
	for _, l := range lines {
		c := parseClaim(l)
		claims = append(claims, &c)
		for i := c.start.x; i < c.end.x; i++ {
			for j := c.start.y; j < c.end.y; j++ {
				key := fmt.Sprintf("%v_%v", i, j)
				claimsAtPoint, _ := points[key]
				claimsAtPoint = append(claimsAtPoint, &c)

				if len(claimsAtPoint) == 2 {
					totalOverlap++
				}
				if len(claimsAtPoint) > 1 {
					markAsOverlap(claimsAtPoint)
				}
				points[key] = claimsAtPoint
			}
		}
	}
	fmt.Println("Total overlap:", totalOverlap)

	for _, c := range claims {
		if !c.overlaps {
			fmt.Println("Non-overlap:", c.id)
			break
		}
	}

}

func parseClaim(s string) claim {
	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	matches := re.FindStringSubmatch(s)

	x := mustAtoi(matches[2])
	y := mustAtoi(matches[3])
	width := mustAtoi(matches[4])
	height := mustAtoi(matches[5])
	return claim{
		id:       mustAtoi(matches[1]),
		start:    position{x, y},
		end:      position{x: x + width, y: y + height},
		size:     size{width, height},
		overlaps: false}
}

func markAsOverlap(claims []*claim) {
	for i := range claims {
		(*claims[i]).overlaps = true
	}
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
