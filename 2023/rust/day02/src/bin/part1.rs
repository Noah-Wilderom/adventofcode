use day02::part1::process;
use std::fs;

fn main() {
    let file = fs::read_to_string("./input1.txt").unwrap();
    println!("{}", process(&file));
}
