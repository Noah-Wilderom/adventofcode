/*
--- Day 3: Gear Ratios ---
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

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

func getMap() []Coords {
	mapList := []Coords{
		{x: -1, y: 0},  // Left
		{x: -1, y: -1}, // Top left
		{x: -1, y: 1},  // Bottom left

		{x: 1, y: 0},  // Right
		{x: 1, y: 1},  // Top right
		{x: 1, y: -1}, // Bottom right
	}

	return mapList
}

func isNumber(input byte) bool {
	_, err := strconv.Atoi(string(input))

	return err == nil
}

func isValid(input byte, nums []*Number, x int, y int) bool {
	for i := range nums {
		for _, coord := range nums[i].coords {
			if coord.x == x && coord.y == y {
				return false
			}
		}
	}

	return true
}

func isSymbol(input byte) bool {

	switch string(input) {
	case "$":
		return true

	case "*":
		return true

	case "#":
		return true

	case "+":
		return true
	}

	return false
}

func calcSum(input []string, nums []*Number) int {
	mapList := getMap()
	sum := 0

	for _, num := range nums {

		maxY := num.coords[len(num.coords)-1].y
		maxX := num.coords[len(num.coords)-1].x
		minY := num.coords[0].y
		minX := num.coords[0].x
		isSummed := false

		for _, m := range mapList {
			maxYIndex := maxY + m.y
			maxXIndex := maxX + m.x
			minYIndex := minY + m.y
			minXIndex := minX + m.x

			if isValidIndex(input, maxYIndex, maxXIndex) {
				if isSymbol(input[maxYIndex][maxXIndex]) {
					fmt.Println(sum, "+", num.num)
					sum += num.num
					isSummed = true
					break
				}

			} else if isValidIndex(input, minYIndex, minXIndex) && isSymbol(input[minYIndex][minXIndex]) {
				fmt.Println(sum, "+", num.num)
				sum += num.num
				isSummed = true
				break
			}
		}

		if !isSummed {
			for _, coords := range num.coords {
				if isValidIndex(input, coords.y+1, coords.x) && isSymbol(input[coords.y+1][coords.x]) {
					fmt.Println(sum, "+", num.num)
					sum += num.num
					isSummed = true
					break
				} else if isValidIndex(input, coords.y-1, coords.x) && isSymbol(input[coords.y-1][coords.x]) {
					fmt.Println(sum, "+", num.num)
					sum += num.num
					isSummed = true
					break
				}
			}
		}

	}

	return sum
}

func isValidIndex(arr []string, row, col int) bool {

	return row >= 0 && row < len(arr) && col >= 0 && col < len(arr[row])
}

func process(input []string) int {
	var nums []*Number

	for y := range input {
		for x := range input[y] {
			if isNumber(input[y][x]) && isValid(input[y][x], nums, x, y) {
				i := 0
				digits := string(input[y][x])
				var coords []*Coords
				coords = append(coords, &Coords{
					x: x,
					y: y,
				})

				for {
					i = i + 1

					if isValidIndex(input, y, x+i) {
						if isNumber(input[y][x+i]) && isValid(input[y][x+i], nums, x+i, y) {
							digits += string(input[y][x+i])
							coords = append(coords, &Coords{
								x: x + i,
								y: y,
							})
						} else {
							break
						}
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

	return calcSum(input, nums)

	// for _, num := range nums {
	// 	var xCoords []string
	// 	var yCoords []string
	// 	for _, coords := range num.coords {
	// 		xCoords = append(xCoords, fmt.Sprint(coords.x))
	// 		yCoords = append(yCoords, fmt.Sprint(coords.y))
	// 	}
	//
	// 	fmt.Println("Num:", num.num, "X:", strings.Join(xCoords, " "), "Y:", strings.Join(yCoords, " "))
	// }
}

func Run() {

	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	fmt.Printf("%d Games\nOutput: %d\n", len(lines), process(lines))
}
