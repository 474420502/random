package random

import (
	"log"
	"testing"
)

func estPhone(t *testing.T) {
	r := New()

	for i := 0; i < 1000; i++ {
		p := r.Extend().Phone()
		log.Println(p.Number)
		log.Println(p.Operators)
	}
}
