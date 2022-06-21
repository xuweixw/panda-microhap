package microhaplotype

import (
	"fmt"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	rdr := Read()

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}
		for i, _ := range rdr.Header.SampleNames {
			//time.Sleep(time.Second)
			var AlellicDepth []int
			if AD, err := variant.GetGenotypeField(variant.Samples[i], "AD", -1); err == nil {
				//	fmt.Printf("AD Type: %T\n", AD)
				AlellicDepth = AD.([]int)
				//	fmt.Printf("AD Type: %T\n", AlellicDepth)
				if ok := IsDiploid(AlellicDepth[0], AlellicDepth[1:]); !ok {
					break
				}
			}
			//fmt.Println(sample)
			if len(rdr.Header.SampleNames) == i+1 {
				// fmt.Println(variant)
				// Segregation
				for i, fam := range Family {

					fatherGT, _ := GetGT(variant, SampleID[fam[0]])
					motherGT, _ := GetGT(variant, SampleID[fam[1]])
					childGT, _ := GetGT(variant, SampleID[fam[2]])
					fatherAD, _ := GetAD(variant, SampleID[fam[0]])
					motherAD, _ := GetAD(variant, SampleID[fam[1]])
					childAD, _ := GetAD(variant, SampleID[fam[2]])
					if ok := Segregate([]Genotype{fatherGT, motherGT}, childGT); !ok {
						fmt.Println(variant.Chrom(), variant.Pos, fatherGT, fatherAD, motherGT, motherAD, childGT, childAD, fam)
						time.Sleep(time.Second)
						break
					}
					if i == len(Family)-1 {
						fmt.Println(variant)
					}
				}
			}
		}
	}
}
