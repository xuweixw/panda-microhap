package main

import (
	. "Microhaplotype/microhaplotype"
	"flag"
	"fmt"
	"os"
)

var (
	MinInterval = flag.Int("minInterval", 3, "specify the minimum interval of a pair of adjacent SNPs")
	MaxInterval = flag.Int("maxInterval", 40, "specify the maximum interval of a pair of adjacent SNPs")
	MaxLen      = flag.Int("maxLen", 110, "specify the maximum length of a microhaplotype")
	Margin      = flag.Int("margin", 100, "specify a margin for PCR primer design")
	Lab         = flag.String("name", "GP", "specify a abbreviation for your laboratory")
	vcfFile     = flag.String("vcf", "", "specify a vcf file")
	_version    = flag.Bool("v", false, "print current version")
	_help       = flag.Bool("h", false, "print help information")
)

const (
	VERSION = "v1.0"
	USAGE   = `	panda-microhap  [options] -vcf in.vcf [ 1> out.vcf 2> out.bed]`
)

func ParseFlag() {
	if len(os.Args) > 1 {
		flag.Parse()
	} else {
		fmt.Println(USAGE)
		flag.Usage()
		os.Exit(0)
	}
	if *_version {
		fmt.Println(VERSION)
		os.Exit(0)
	} else if *vcfFile == "" || *_help {
		fmt.Println(USAGE)
		flag.Usage()
		os.Exit(0)
	}
}
func main() {
	ParseFlag()
	//"data/microhapl-trialellic.vcf"
	rdr := Read(*vcfFile)
	var spr = NewSpring(uint(*MaxInterval), uint(*MinInterval), uint(*MaxLen))
	MHnamed := Nomenclature(*Lab)
	for {
		variant := rdr.Read()
		if variant == nil {
			outVCF, outBED := String(spr.PopAll(), MHnamed, uint64(*Margin))
			fmt.Fprint(os.Stdout, outVCF)
			fmt.Fprint(os.Stderr, outBED)
			break
		}
		//fmt.Println(*variant)
		if r := spr.Add(variant); r != nil {
			outVCF, outBED := String(r, MHnamed, uint64(*Margin))
			fmt.Print()
			fmt.Fprint(os.Stdout, outVCF)
			fmt.Fprint(os.Stderr, outBED)
		}
	}
}
