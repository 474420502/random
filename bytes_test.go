package random

import (
	"testing"
)

func TestBytes(t *testing.T) {
	r := New()
	for i := 0; i < 1000; i++ {
		b := r.Bytes(0, 100)
		if len(b) > 100 {
			panic("bytes len is error len(b) >= 100")
		}
	}

	for i := 0; i < 1000; i++ {
		b := r.Bytes(50, 200)
		if len(b) > 200 {
			panic("bytes len is error len(b) >= 200 ")
		}

		if len(b) < 50 {
			panic("bytes len is error  len(b) < 50 ")
		}
	}
}
