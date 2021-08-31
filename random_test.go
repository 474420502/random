package random

import (
	"encoding/gob"
	"log"
	"os"
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestGetData(t *testing.T) {
	r := New(1630407844130142968)
	Use(DataNameChinese)
	for i := 0; i < 100; i++ {
		log.Println(r.Extend().FullName())
	}

}

func WriteCase1(t *testing.T) {
	f, err := os.OpenFile("firstnames.gob.zst", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}

	enc, err := zstd.NewWriter(f)
	if err != nil {
		panic(err)
	}

	genc := gob.NewEncoder(enc)
	if err != nil {
		panic(err)
	}

	err = genc.Encode(firstNameData)
	if err != nil {
		panic(err)
	}

	enc.Close()
	f.Close()
}

func TestCase2(t *testing.T) {
	// zstd.NewWriter()
	f, err := os.Open("firstnames.gob")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(f)
	var lastnames map[string]int
	err = dec.Decode(&lastnames)
	if err != nil {
		panic(err)
	}

	// data, err := ioutil.ReadAll(f)
	// if err != nil {
	// 	panic(err)
	// }

	f, err = os.OpenFile("lastnames.gob.zstd", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}

	enc, err := zstd.NewWriter(f)
	if err != nil {
		panic(err)
	}

	gob.NewEncoder(enc).Encode(lastnames)

	// _, err = enc.Write(data)
	// if err != nil {
	// 	panic(err)
	// }

}
