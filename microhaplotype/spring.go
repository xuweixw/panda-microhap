package microhaplotype

import (
	"errors"
	"fmt"
	"github.com/brentp/vcfgo"
	"log"
	"strconv"
	"strings"
)

type Spring struct {
	maxInterval, minInterval uint
	//margin                   uint // margin for PCR primer design
	data      []*vcfgo.Variant
	maxLength uint
}

func NewSpring(maxIntvl, minIntvl, maxLen uint) *Spring {
	return &Spring{maxInterval: maxIntvl,
		minInterval: minIntvl,
		data:        make([]*vcfgo.Variant, 0),
		maxLength:   maxLen,
	}
}

// Helix return that this Spring has contained how many elements.
func (s *Spring) Helix() int {
	return len(s.data)
}

func (s *Spring) Len() int {
	if len(s.data) <= 1 {
		return 0
	} else {
		return int(s.data[len(s.data)-1].Pos - s.data[0].Pos)
	}
}

func (s *Spring) Top() *vcfgo.Variant {
	return s.data[len(s.data)-1]
}

// PopAll evaluates whether is a microhaptype in current Spring。
// If there is only an element or the length of Spring exceed maxLength， clear all data and return nil.
// If Spring element contains a Microhaplotype, return this as a slice.
func (s *Spring) PopAll() []*vcfgo.Variant {
	if len(s.data) == 1 || s.maxLength < uint(s.Len()) {
		// fmt.Println("here-----", len(s.data) == 1, s.maxLength, uint(s.Len()))
		s.data = make([]*vcfgo.Variant, 0)
		return nil
	} else if !CheckMH(s.data) {
		return nil
	} else {
		return s.data
	}
}

func (s *Spring) Add(variant *vcfgo.Variant) []*vcfgo.Variant {
	//fmt.Println(*s, "Add-----")
	if s.Helix() == 0 {
		//fmt.Println("1111111")
		s.data = append(s.data, variant)
		return nil
	} else if variant.Chromosome != s.Top().Chromosome ||
		variant.Pos-s.Top().Pos > uint64(s.maxInterval) ||
		variant.Pos-s.Top().Pos < uint64(s.minInterval) {
		//fmt.Println(variant.Pos-s.Top().Pos, uint64(s.minInterval))
		//fmt.Println("22222222")
		var res = s.PopAll()
		s.data = []*vcfgo.Variant{variant}
		return res
	} else {
		//fmt.Println("33333333")
		s.data = append(s.data, variant)
		return nil
	}
}

func String(variant []*vcfgo.Variant, named func(*vcfgo.Variant) string, Margin uint64) (string, string) {
	if len(variant) == 0 || IsInDel(variant) {
		return "", ""
	}
	var s strings.Builder
	//s.WriteString(fmt.Sprintln("Microhaplotype"))
	name := named(variant[0])
	mhSNP := len(variant)
	//
	for i := range variant {
		variant[i].Info().Set("MH_ID", name)
		variant[i].Info().Set("MH_SNP", mhSNP)
		s.WriteString(variant[i].String() + "\n")
	}
	Range := fmt.Sprintf("%s\t%d\t%d\t%s\n",
		variant[0].Chromosome,
		variant[0].Pos-Margin,
		variant[len(variant)-1].Pos+Margin,
		name)
	return s.String(), Range
}

func IsInDel(variant []*vcfgo.Variant) bool {
	for i := range variant {
		for _, alt := range variant[i].Alt() {
			if alt == "*" {
				return true
			}
		}
	}
	return false
}

// Nomenclature builds a named system for Microhaps, eg, "mh01GP-000001"
// The "mh" refers to MicroHaplotype, the "01" or "0X" refers to the numbering of chromosome,
// the "GP" refers to the abbreviation of Panda Genetic Resources lab.
func Nomenclature(lab string) func(variant *vcfgo.Variant) string {
	var (
		MarkerType    = "mh"
		InitialNumber = 1
	)
	return func(variant *vcfgo.Variant) string {
		var chr = variant.Chromosome
		if len(chr) == 1 {
			chr = "0" + chr
		} else if len(chr) > 2 {
			log.Panicln("Chromosome name must have 2-digit，but get: ", chr, " in line ", variant.LineNumber)
		}
		var name = fmt.Sprintf("%s%s%s-%06d", MarkerType, chr, lab, InitialNumber)
		InitialNumber++
		return name
	}
}

// CheckMH checks all variants in a microhap whichever law diploid, segregation, genotype polymorphsim.
func CheckMH(variants []*vcfgo.Variant) bool {
	for _, variant := range variants {
		// Check Diploid
		for _, name := range variant.Header.SampleNames {
			ADslice, _ := GetAD(variant, name)
			if !IsDiploid(ADslice[0], ADslice[1:]) {
				return false
			}
		}
		// Check Segregation
		for _, fam := range Family {
			fatherGT, _ := GetGT(variant, SampleID[fam[0]])
			motherGT, _ := GetGT(variant, SampleID[fam[1]])
			childGT, _ := GetGT(variant, SampleID[fam[2]])
			if ok := Segregate([]Genotype{fatherGT, motherGT}, childGT); !ok {
				//fmt.Println(variant.Chrom(), variant.Pos, fatherGT, fatherAD, motherGT, motherAD, childGT, childAD, fam)
				//time.Sleep(time.Second)
				return false
			}
		}
		// Check genotype polymophsim
		sampleNames := variant.Header.SampleNames
		firstGT, _ := GetGT(variant, sampleNames[0])
		for i := 1; i < len(sampleNames); i++ {
			nextGT, _ := GetGT(variant, sampleNames[i])
			if !firstGT.Equal(nextGT) {
				return true
			} else if i == len(sampleNames)-1 {
				return false
			}
		}
	}
	return true
}

func GetGT(variant *vcfgo.Variant, sample string) (Genotype, error) {
	for i, s := range variant.Header.SampleNames {
		if s == sample {

			if GT, err := variant.GetGenotypeField(variant.Samples[i], "GT", -1); err == nil {
				GTString := GT.(string)
				diploid := strings.Split(GTString, "/")
				ploidA, _ := strconv.ParseUint(diploid[0], 10, 64)
				ploidB, _ := strconv.ParseUint(diploid[1], 10, 64)
				genotype := Genotype{int(ploidA), int(ploidB)}
				return genotype, nil
			}
		}
	}
	var err = errors.New("New error")
	return Genotype{0, 0}, err
}

func GetAD(variant *vcfgo.Variant, sample string) ([]int, error) {
	for i, s := range variant.Header.SampleNames {
		if s == sample {
			if AD, err := variant.GetGenotypeField(variant.Samples[i], "AD", -1); err == nil {
				ADslice := AD.([]int)
				return ADslice, nil
			}
		}
	}
	var err = errors.New("New error")
	return []int{}, err
}
