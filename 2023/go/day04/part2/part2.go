package part2

import (
	"fmt"
	"os"
	"strings"
)

func process(input []string) int {
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
