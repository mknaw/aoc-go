package common

import "fmt"

type ProblemPart int

const (
	A ProblemPart = iota
	B
)

func (p ProblemPart) String() string {
	switch p {
	case A:
		return "A"
	case B:
		return "B"
	default:
		return "Unknown"
	}
}

type Problem struct {
	Year int
	Day  int
	Foo  int
	Part ProblemPart
}

func (p Problem) String() string {
	return fmt.Sprintf("Problem %d-%02d-%s", p.Year, p.Day, p.Part.String())
}

type TestCase struct {
	Input    string
	Expected int
}
