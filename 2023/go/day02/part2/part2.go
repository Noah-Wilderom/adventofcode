package part2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_hands(game string) []string {
	game = strings.Split(game, ":")[1]
	hands := strings.Split(game, ";")

	return hands
}

func parse_color(hand string, color string) (int, error) {

	colors := strings.Split(hand, ",")

	for _, v := range colors {
		if strings.Contains(v, color) {
			intVal := strings.Split(v, color)[0]
			intVal = strings.Replace(intVal, " ", "", -1)
			return strconv.Atoi(intVal)
		}
	}

	return 0, nil
}

func sum_game_hands(hands []string) int {
	green := 0
	red := 0
	blue := 0

	for _, hand := range hands {
		greenParsed, err := parse_color(hand, "green")
		if err != nil {
			panic(err)
		}
		redParsed, err := parse_color(hand, "red")
		if err != nil {
			panic(err)
		}
		blueParsed, err := parse_color(hand, "blue")
		if err != nil {
			panic(err)
		}

		if greenParsed > green {
			green = greenParsed
		}
		if redParsed > red {
			red = redParsed
		}
		if blueParsed > blue {
			blue = blueParsed
		}
	}

	return green * red * blue
}

func process(input []string) int {
	sum := 0

	for _, game := range input {
		if game == "" {
			continue
		}
		hands := parse_hands(game)

		sum = sum + sum_game_hands(hands)
	}

	return sum
}

func Run() {

	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	fmt.Printf("%d Games\nOutput: %d\n", len(lines), process(lines))
}
