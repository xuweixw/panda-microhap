package microhaplotype

import (
	"fmt"
	"testing"
)

func TestNewSpring(t *testing.T) {
	rdr := Read()
	var spr = NewSpring(40, 3, 110)
	MHnamed := Nomenclature("GP")
	for {
		variant := rdr.Read()
		if variant == nil {
			fmt.Print(String(spr.PopAll(), MHnamed))
			break
		}
		//fmt.Println(*variant)
		if r := spr.Add(variant); r != nil {
			fmt.Print(String(r, MHnamed))
		}
	}
}

func TestNomenclature(t *testing.T) {
	// the line below must be added in vcf file.
	// ##INFO=<ID=MH,Number=1,Type=String,Description="this variant and its relatives make a microhap">
	rdr := Read()

	MHnamed := Nomenclature("GP")
	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}
		variant.Info().Set("MH", MHnamed(variant))
		fmt.Println(variant)
	}
	// Output:
	//mh0XGP-000001
	//mh09GP-000003
}
