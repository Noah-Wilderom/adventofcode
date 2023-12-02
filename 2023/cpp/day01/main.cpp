#include "../shared/Test.hpp"

int process(std::string input)
{
	//
}

int main() {
	AdventOfCode::Test *test;

	std::string input = "1abc2"
		"pqr3stu8vwx"
		"a1b2c3d4e5f"
		"treb7uchet";

	test->expectEqual<int, int>(process(input), 142);
}
