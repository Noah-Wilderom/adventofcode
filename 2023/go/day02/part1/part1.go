package part1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_BLUE  = 14
	MAX_GREEN = 13
	MAX_RED   = 12
)

func parse_game(game string) (int, error) {
	split := strings.Split(game, ":")

	gameId := strings.Split(split[0], "Game ")

	return strconv.Atoi(gameId[1])
}

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

func is_valid_game(hands []string) bool {

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

		if greenParsed > MAX_GREEN || redParsed > MAX_RED || blueParsed > MAX_BLUE {
			return false
		}
	}
	return true
}

func process(input []string) int {
	games := 0

	for _, game := range input {
		if game == "" {
			continue
		}

		gameId, err := parse_game(game)
		if err != nil {
			panic(err)
		}

		hands := parse_hands(game)

		if is_valid_game(hands) {
			games = games + gameId
		}
	}

	return games
}

func Run() {

	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	fmt.Printf("%d Games\nOutput: %d\n", len(lines), process(lines))
}
