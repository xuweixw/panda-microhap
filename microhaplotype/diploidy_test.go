package microhaplotype

import (
	"fmt"
	"testing"
)

func TestIsDiploid(t *testing.T) {
	var (
		Test = [][]int{
			[]int{23, 54},
			[]int{0, 154},
			[]int{42, 83, 0},
			[]int{38, 64, 12},
			[]int{0, 106, 0},
			[]int{23, 54, 45, 0},
			[]int{29, 0, 35, 0},
			[]int{34, 2, 0, 34},
		}
		Result = []bool{true, true, true, false, true, false, true, false}
		s      = NewSummary()
	)
	for i := 0; i < len(Test); i++ {
		if ok := s.Check(Test[i][0], Test[i][1:]); ok != Result[i] {
			t.Fatalf("In %v, Except %v, Given %v\n", Test[i], ok, Result[i])
		}
	}
	fmt.Println(s.success, s.failure)
}
