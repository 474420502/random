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

func BenchmarkIndustry(b *testing.B) {

	r := New()
	Use(DataIndustryChina)

	for i := 0; i < b.N; i++ {
		r.Extend().Industry()
	}
}

func BenchmarkR(b *testing.B) {

}
