package microhaplotype

import (
	"testing"
)

func TestSegregate(t *testing.T) {
	var (
		Parents = [][]Genotype{
			{[2]byte{0, 0}, [2]byte{1, 1}, [2]byte{0, 1}}, // P: "0/0", "1/1"; Child: "0/1", T
			{[2]byte{0, 1}, [2]byte{0, 1}, [2]byte{0, 1}}, // P: "0/1", "0/1"; Child: "0/1", T
			{[2]byte{0, 2}, [2]byte{1, 2}, [2]byte{2, 2}}, // P: "0/2", "1/2"; Child: "2/2", T
			{[2]byte{2, 2}, [2]byte{0, 2}, [2]byte{1, 2}}, // P: "2/2", "0/2"; Child: "1/2", F
			{[2]byte{1, 1}, [2]byte{0, 2}, [2]byte{0, 1}}, // P: "1/1", "0/2"; Child: "0/1", T
			{[2]byte{1, 1}, [2]byte{0, 2}, [2]byte{0, 1}}, // P: "1/1", "0/2"; Child: "0/1", T
		}
		expects = []bool{true, true, true, false, true, true}
	)

	for i, P := range Parents {
		if get := Segregate(P[:2], P[2]); get != expects[i] {
			t.Errorf("In %v - > %v, expect %v, get %v", P[:2], P[2], expects[i], get)
		}
	}
}
