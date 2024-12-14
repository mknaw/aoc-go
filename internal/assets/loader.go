package assets

import (
	"aoc-go/common"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LoadProblem(problem *common.Problem) (string, error) {
	return readFile(problem, false)
}

func LoadTests(problem *common.Problem) ([]common.TestCase, error) {
	input, err := readFile(problem, true)
	if err != nil {
		// Fine if there's no test file, just return empty `[]TestCase` and we won't test.
		input = ""
	}
	return splitTestCases([]byte(input))
}

func readFile(problem *common.Problem, isTest bool) (string, error) {
	// Get the project root directory (where go.mod is)
	projectRoot, err := findProjectRoot()
	if err != nil {
		return "", err
	}

	var filename string
	if isTest {
		filename = fmt.Sprintf(
			"tests/%d/%02d%s.txt",
			problem.Year,
			problem.Day,
			strings.ToLower(problem.Part.String()),
		)
	} else {
		filename = fmt.Sprintf("problems/%d/%02d.txt", problem.Year, problem.Day)
	}
	assetPath := filepath.Join(projectRoot, "assets", filename)
	bytes, err := os.ReadFile(assetPath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func splitTestCases(data []byte) ([]common.TestCase, error) {
	var result []common.TestCase
	start := 0

	i := 0
	for i < len(data)-3 {
		if bytes.HasPrefix(data[i:], []byte(">>> ")) {
			input := string(data[start:i])
			i += 4

			j := i
			for j < len(data) && data[j] >= '0' && data[j] <= '9' {
				j++
			}
			expected, err := strconv.Atoi(string(data[i:j]))
			if err != nil {
				return nil, err
			}

			i = j
			for i < len(data) && data[i] == '\n' {
				i += 1
			}
			start = i

			result = append(result, common.TestCase{Input: input, Expected: expected})
		}
		i++
	}

	return result, nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
