package random

import (
	"log"
	"testing"
)

func estText(t *testing.T) {
	// var s = rune('㐀')
	// var e = rune('鿯')
	r := New()
	for i := 0; i < 100; i++ {
		log.Println(r.Extend().TextRangeN(2, 2))
	}

}
