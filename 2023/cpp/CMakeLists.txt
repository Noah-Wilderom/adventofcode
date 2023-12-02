cmake_minimum_required(VERSION 3.12)

# Check if the day directory is provided as an argument
if(NOT DEFINED DAY_DIRECTORY)
    message(FATAL_ERROR "Please provide the day directory as an argument. Example: cmake -DDAY_DIRECTORY=01")
endif()

project(AdventOfCode)

set(CMAKE_CXX_STANDARD 14)

# Include directories
include_directories(${CMAKE_SOURCE_DIR}/shared)

# Add the shared library
add_library(AdventOfCode_shared INTERFACE)
target_include_directories(AdventOfCode_shared INTERFACE ${CMAKE_SOURCE_DIR}/shared)

# Add the day executable
add_executable(day-${DAY_DIRECTORY}
    ${CMAKE_SOURCE_DIR}/day${DAY_DIRECTORY}/main.cpp
)

# Link the shared library to the day executable
target_link_libraries(day-${DAY_DIRECTORY}
	PUBLIC AdventOfCode_shared
)
