package y2024

import (
	"aoc-go/common/parsing"
	"strings"
)

func Day02(input string, removes int) (int, error) {
	safeCount := 0
	for _, row := range strings.Split(input, "\n") {
		ints, err := parsing.RowToInts(row)
		if err != nil {
			return 0, err
		}

		if len(ints) > 0 && isSafe(ints, removes) {
			safeCount++
		}
	}

	return safeCount, nil
}

func Day02a(input string) (int, error) {
	return Day02(input, 0)
}

func Day02b(input string) (int, error) {
	return Day02(input, 1)
}

func isDeltaOk(x int, y int, isAsc bool) bool {
	delta := y - x
	if !isAsc {
		delta *= -1
	}
	return 0 < delta && delta < 4
}

func isSafe(ints []int, removes int) bool {
	return isSafeOneWay(ints, removes, true) || isSafeOneWay(ints, removes, false)
}

func isSafeOneWay(ints []int, removes int, isAsc bool) bool {
	for i := 1; i < len(ints); i++ {
		if isDeltaOk(ints[i-1], ints[i], isAsc) {
			continue
		}
		if removes > 0 {
			if i == len(ints)-1 {
				// We're at the end, so it works if we remove the last one.
				return true
			} else if i < len(ints)-1 && isDeltaOk(ints[i-1], ints[i+1], isAsc) {
				i++
				removes--
				continue
			} else if i == 1 {
				// Might have to remove the first element. We'll see next iteration.
				removes--
				continue
			}
		}
		return false
	}
	return true
}
