package part1

import "testing"

func TestAlgo(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	output := process(lines)

	if output != 4361 {
		t.Errorf("\nExpected output: 4361\nActual output: %d\n", output)
	}
}
