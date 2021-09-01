package random

import (
	"testing"
)

func BenchmarkChars(b *testing.B) {

	r := New()
	Use(DataEmailChina)

	for i := 0; i < b.N; i++ {
		r.Extend().Email()
	}
}

func BenchmarkR(b *testing.B) {

}
