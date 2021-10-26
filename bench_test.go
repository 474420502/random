package random

import (
	"log"
	"testing"
)

func init() {
	// ConfigSourcePrefix("https://gitee.com/githubcontent/random_data/raw/master/")
}

func BenchmarkEmail(b *testing.B) {
	r := New()
	Use(DataEmailChina)
	for i := 0; i < b.N; i++ {
		r.Extend().Email()
	}
}

func BenchmarkIndustry(b *testing.B) {
	b.StopTimer()
	r := New()
	Use(DataIndustryChina)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		r.Extend().Industry()
	}
}

func BenchmarkName(b *testing.B) {

	b.StopTimer()
	r := New()
	Use(DataNameChina)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Extend().FullName()
	}
}

func BenchmarkProvince(b *testing.B) {
	b.StopTimer()
	r := New()
	Use(DataProvinceChina)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Extend().FullName()
	}
}

func BenchmarkCity(b *testing.B) {
	b.StopTimer()
	r := New()
	Use(DataCityChina)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Extend().City()
	}
}

func BenchmarkArea(b *testing.B) {
	b.StopTimer()
	r := New()
	Use(DataAreaChina)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Extend().Area()
	}
}

func BenchmarkTown(b *testing.B) {
	b.StopTimer()
	r := New()
	Use(DataTownStreetChina)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Extend().Town()
	}
}

func BenchmarkCountry(b *testing.B) {
	b.StopTimer()
	r := New()
	Use(DataCountryChina)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r.Extend().Country()
	}
}

func BenchmarkBytes(b *testing.B) {
	b.StopTimer()
	r := New()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Println(r.Bytes(0, 100))
	}
}
