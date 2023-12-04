package part1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	num    int
	coords []*Coords
}

type Coords struct {
	x int
	y int
}

func isNumber(input byte) bool {
	_, err := strconv.Atoi(string(input))

	return err == nil
}

func process(input []string) int {
	var nums []*Number

	for y := range input {
		for x := range input[y] {
			if isNumber(input[y][x]) {
				i := 0
				digits := string(input[y][x])
				var coords []*Coords
				coords = append(coords, &Coords{
					x: x,
					y: y,
				})

				for {
					i = i + 1

					if isNumber(input[y][x+i]) {
						digits += string(input[y][x+i])
						coords = append(coords, &Coords{
							x: x + i,
							y: y,
						})
					} else {
						break
					}
				}

				numeralDigits, err := strconv.Atoi(digits)
				if err != nil {
					panic(err)
				}

				nums = append(nums, &Number{
					num:    numeralDigits,
					coords: coords,
				})

			}

		}
	}

	for _, num := range nums {
		var xCoords []string
		var yCoords []string
		for _, coords := range num.coords {
			xCoords = append(xCoords, fmt.Sprint(coords.x))
			yCoords = append(yCoords, fmt.Sprint(coords.y))
		}

		fmt.Println("Num:", num.num, "X:", strings.Join(xCoords, " "), "Y:", strings.Join(yCoords, " "))
	}

	return 1
}

func Run() {

	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	fmt.Printf("%d Games\nOutput: %d\n", len(lines), process(lines))
}
