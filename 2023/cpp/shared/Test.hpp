#ifndef TEST_H
#define TEST_H

#include <iostream>
#include <string>

namespace AdventOfCode {
struct Test {

  template <typename T1, typename T2> void expectEqual(T1 item1, T2 item2) {
    if (item1 == item2) {
      std::cout << "Test Success!" << std::endl;
    } else {
      std::cout << "Test Failed: " << item1 << " does not equal " << item2
                << std::endl;
    }

    std::cout << std::endl
              << "Left: " << item1 << std::endl
              << "Right: " << item2 << std::endl;
  }
};
}; // namespace AdventOfCode

#endif // !TEST_H
