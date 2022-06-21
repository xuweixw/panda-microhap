package microhaplotype

// IsDiploid checks whether the AD (allele depth) in Sample field of vcf-format represents diploidy and
// applies principally to tri-allele or quad-allele.
func IsDiploid(ref int, alts []int) bool {
	// only two numbers are more than zero.
	var moreThanZero int
	if ref > 0 {
		moreThanZero++
	}
	for i := range alts {
		if alts[i] > 0 {
			moreThanZero++
		}
	}
	if moreThanZero <= 2 {
		return true
	}
	return false
}

func check() func(int, []int) bool {
	return IsDiploid
}

type Summary struct {
	success, failure int
	handle           func(int, []int) bool
}

func NewSummary() *Summary {
	return &Summary{0, 0, IsDiploid}
}

func (s *Summary) Check(ref int, alts []int) bool {
	if ok := s.handle(ref, alts); ok {
		s.success++
		return ok
	} else {
		s.failure++
		return ok
	}
}
