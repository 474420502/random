package random

import (
	"log"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	r := New()
	now := time.Now().Truncate(24 * time.Hour)
	for i := 0; i < 100; i++ {
		log.Println(r.Extend().DateRange(now.Add(-time.Hour*24), now).UTC())
	}
}

func TestDateRange(t *testing.T) {
	r := New()

	for i := 0; i < 100; i++ {
		log.Println(r.Extend().DateRange("2021-07-01", "2021-07-02").UTC())
	}
}
