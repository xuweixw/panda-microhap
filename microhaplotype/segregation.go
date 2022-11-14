package microhaplotype

import "log"

// law of segregation
// Gregor Johann Mendel

type Genotype [2]byte // "0/0, 0/1, 1/1, 0/2, 1/2, 2/2, ./."

func (g Genotype) Equal(genotype Genotype) bool {
	return g[0] == genotype[0] && g[1] == genotype[1]
}

func Segregate(parents []Genotype, children Genotype) bool {
	if len(parents) != 2 {
		log.Panicln("Must have two parents")
		return false
	}
	var expect = []Genotype{
		{parents[0][0], parents[1][0]},
		{parents[0][0], parents[1][1]},
		{parents[0][1], parents[1][0]},
		{parents[0][1], parents[1][1]},
	}
	for i := range expect {
		// 基因型分离时，可能是大数在前，需要考虑这种情况
		if expect[i].Equal(children) || expect[i].Equal([2]byte{children[1], children[0]}) {
			break
		} else if i == 3 {
			return false
		}
	}
	return true
}
