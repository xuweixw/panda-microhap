package microhaplotype

import (
	"github.com/brentp/vcfgo"
	"log"
	"os"
)

func Read(path string) *vcfgo.Reader {
	f, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	rdr, err := vcfgo.NewReader(f, false)
	if err != nil {
		log.Panicln(err)
	}
	return rdr
}
