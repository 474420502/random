package random

import (
	"encoding/gob"
	"log"
	"os"
	"testing"

	"github.com/474420502/gcurl"
	"github.com/klauspost/compress/zstd"
	"github.com/tidwall/gjson"
)

func TestGetData(t *testing.T) {
	Use(DataCountryChina)
	r := New(1630407844130142968)
	for i := 0; i < 100; i++ {
		log.Println(r.Extend().Country().LocalName)
	}

}

var data []*Country

func TestWriteCase1(t *testing.T) {

	resp, err := gcurl.Parse("https://raw.githubusercontent.com/uiwjs/province-city-china/gh-pages/country.json").Temporary().Execute()
	if err != nil {
		panic(err)
	}
	for _, g := range gjson.Parse(string(resp.Content())).Array() {
		country := &Country{}
		country.ID = int(g.Get("id").Int())
		country.LocalName = g.Get("cnname").String()
		country.Name = g.Get("name").String()
		country.FullName = g.Get("fullname").String()
		country.Alpha2 = g.Get("alpha2").String()
		country.Alpha3 = g.Get("alpha3").String()
		country.Code = int(g.Get("numeric").Int())

		data = append(data, country)
	}

	f, err := os.OpenFile("country.gob.zst", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
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

	err = genc.Encode(data)
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
