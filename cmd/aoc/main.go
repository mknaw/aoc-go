package main

import (
	"aoc-go/common"
	"aoc-go/internal/assets"
	"aoc-go/internal/solutions/y2024"
	"flag"
	"fmt"
	"time"
)

func main() {

	year := flag.Int("year", 2024, "Year of puzzle")
	day := flag.Int("day", 1, "Day of puzzle")

	flag.Parse()

	solutionMap := map[common.Problem]func(string) (int, error){
		{Year: 2024, Day: 1, Part: common.A}:  y2024.Day01a,
		{Year: 2024, Day: 1, Part: common.B}:  y2024.Day01b,
		{Year: 2024, Day: 2, Part: common.A}:  y2024.Day02a,
		{Year: 2024, Day: 2, Part: common.B}:  y2024.Day02b,
		{Year: 2024, Day: 3, Part: common.A}:  y2024.Day03a,
		{Year: 2024, Day: 3, Part: common.B}:  y2024.Day03b,
		{Year: 2024, Day: 4, Part: common.A}:  y2024.Day04a,
		{Year: 2024, Day: 4, Part: common.B}:  y2024.Day04b,
		{Year: 2024, Day: 5, Part: common.A}:  y2024.Day05a,
		{Year: 2024, Day: 5, Part: common.B}:  y2024.Day05b,
		{Year: 2024, Day: 6, Part: common.A}:  y2024.Day06a,
		{Year: 2024, Day: 6, Part: common.B}:  y2024.Day06b,
		{Year: 2024, Day: 7, Part: common.A}:  y2024.Day07a,
		{Year: 2024, Day: 7, Part: common.B}:  y2024.Day07b,
		{Year: 2024, Day: 8, Part: common.A}:  y2024.Day08a,
		{Year: 2024, Day: 8, Part: common.B}:  y2024.Day08b,
		{Year: 2024, Day: 9, Part: common.A}:  y2024.Day09a,
		{Year: 2024, Day: 9, Part: common.B}:  y2024.Day09b,
		{Year: 2024, Day: 10, Part: common.A}: y2024.Day10a,
		{Year: 2024, Day: 10, Part: common.B}: y2024.Day10b,
		{Year: 2024, Day: 11, Part: common.A}: y2024.Day11a,
		{Year: 2024, Day: 11, Part: common.B}: y2024.Day11b,
		{Year: 2024, Day: 12, Part: common.A}: y2024.Day12a,
		{Year: 2024, Day: 12, Part: common.B}: y2024.Day12b,
		{Year: 2024, Day: 13, Part: common.A}: y2024.Day13a,
		{Year: 2024, Day: 13, Part: common.B}: y2024.Day13b,
		{Year: 2024, Day: 14, Part: common.A}: y2024.Day14a,
		{Year: 2024, Day: 14, Part: common.B}: y2024.Day14b,
		{Year: 2024, Day: 15, Part: common.A}: y2024.Day15a,
		{Year: 2024, Day: 15, Part: common.B}: y2024.Day15b,
		{Year: 2024, Day: 16, Part: common.A}: y2024.Day16a,
		{Year: 2024, Day: 16, Part: common.B}: y2024.Day16b,
		{Year: 2024, Day: 17, Part: common.A}: y2024.Day17a,
		{Year: 2024, Day: 17, Part: common.B}: y2024.Day17b,
		{Year: 2024, Day: 18, Part: common.A}: y2024.Day18a,
		{Year: 2024, Day: 18, Part: common.B}: y2024.Day18b,
		{Year: 2024, Day: 19, Part: common.A}: y2024.Day19a,
		{Year: 2024, Day: 19, Part: common.B}: y2024.Day19b,
		{Year: 2024, Day: 20, Part: common.A}: y2024.Day20a,
		{Year: 2024, Day: 20, Part: common.B}: y2024.Day20b,
		{Year: 2024, Day: 20, Part: common.A}: y2024.Day20a,
		{Year: 2024, Day: 20, Part: common.B}: y2024.Day20b,
		{Year: 2024, Day: 21, Part: common.A}: y2024.Day21a,
		{Year: 2024, Day: 21, Part: common.B}: y2024.Day21b,
		{Year: 2024, Day: 22, Part: common.A}: y2024.Day22a,
		{Year: 2024, Day: 22, Part: common.B}: y2024.Day22b,
		{Year: 2024, Day: 23, Part: common.A}: y2024.Day23a,
		{Year: 2024, Day: 23, Part: common.B}: y2024.Day23b,
		{Year: 2024, Day: 24, Part: common.A}: y2024.Day24a,
		{Year: 2024, Day: 24, Part: common.B}: y2024.Day24b,
		{Year: 2024, Day: 25, Part: common.A}: y2024.Day25a,
		{Year: 2024, Day: 25, Part: common.B}: y2024.Day25b,
		{Year: 2024, Day: 26, Part: common.A}: y2024.Day26a,
		{Year: 2024, Day: 26, Part: common.B}: y2024.Day26b,
		{Year: 2024, Day: 27, Part: common.A}: y2024.Day27a,
		{Year: 2024, Day: 27, Part: common.B}: y2024.Day27b,
		{Year: 2024, Day: 28, Part: common.A}: y2024.Day28a,
		{Year: 2024, Day: 28, Part: common.B}: y2024.Day28b,
		{Year: 2024, Day: 29, Part: common.A}: y2024.Day29a,
		{Year: 2024, Day: 29, Part: common.B}: y2024.Day29b,
		{Year: 2024, Day: 30, Part: common.A}: y2024.Day30a,
		{Year: 2024, Day: 30, Part: common.B}: y2024.Day30b,
		{Year: 2024, Day: 31, Part: common.A}: y2024.Day31a,
		{Year: 2024, Day: 31, Part: common.B}: y2024.Day31b,
	}

	for _, problemPart := range []common.ProblemPart{common.A, common.B} {
		problem := common.Problem{
			Year: *year,
			Day:  *day,
			Part: problemPart,
		}

		fmt.Printf("Attempting %s ...\n", problem)

		solutionFunc, ok := solutionMap[problem]
		if !ok {
			fmt.Println("Solution not found!")
			continue
		}

		tests, err := assets.LoadTests(&problem)
		if err != nil {
			fmt.Println(err)
			return
		}
		testsPass := true
		for _, test := range tests {
			res, err := solutionFunc(test.Input)
			if err != nil {
				fmt.Println(err)
				testsPass = false
			}
			if testsPass && res != test.Expected {
				fmt.Printf("Expected %d, got %d\n", test.Expected, res)
				testsPass = false
			}
		}
		if testsPass {
			if len(tests) > 0 {
				fmt.Printf("%d tests passed\n", len(tests))
			}
		} else {
			fmt.Println("Tests failed :(")
			return
		}

		input, err := assets.LoadProblem(&problem)
		if err != nil {
			fmt.Println(err)
			return
		}

		start := time.Now()
		res, err := solutionFunc(input)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("-> %d\n%s\n\n", res, elapsed)

	}
}
