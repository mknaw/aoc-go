package parsing

import (
	"strconv"
	"strings"
)

func RowToInts(row string) ([]int, error) {
	// Given a string of integers separated by spaces, return a slice of integers.
	ints := []int{}
	for _, field := range strings.Fields(row) {
		i, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
