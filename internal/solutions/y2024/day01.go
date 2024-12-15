package y2024

import (
	ds "aoc-go/common/datastructures"
	"aoc-go/common/parsing"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day01a(input string) (int, error) {
	left := []int{}
	right := []int{}

	for _, row := range strings.Split(input, "\n") {

		ints, err := parsing.RowToInts(row)
		if err != nil {
			return 0, err
		}
		if len(ints) != 2 {
			continue
		}

		left = append(left, ints[0])
		right = append(right, ints[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	res := 0
	for i := range left {
		if right[i] > left[i] {
			res += right[i] - left[i]
		} else {
			res += left[i] - right[i]
		}
	}

	return res, nil
}

func Day01b(input string) (int, error) {
	left := []int{}
	right := ds.NewCounter[int]()

	for _, row := range strings.Split(input, "\n") {

		fields := strings.Fields(row)
		if len(fields) != 2 {
			continue
		}
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Println("Error converting to int:", fields[0])
			return 0, err
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("Error converting to int:", fields[1])
			return 0, err
		}

		left = append(left, x)
		right.Inc(y)
	}

	sort.Ints(left)

	res := 0
	for _, k := range left {
		res += k * right.Get(k)
	}

	return res, nil
}
