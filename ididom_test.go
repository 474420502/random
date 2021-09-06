package random

import (
	"log"
	"testing"

	"github.com/474420502/extractor"
	"github.com/474420502/requests"
)

func TestIdidom(t *testing.T) {
	r := New()
	Use(DataIdidomChina)
	for i := 0; i < 1000; i++ {
		log.Println(r.Extend().Ididom())
	}
}

func estIdidom(t *testing.T) {

	u := "https://www.y5000.com/idiom/list-0-0-5.html"
	tp := requests.NewSession().Get(u)
	p := tp.PathParam("list-0-0-(\\d+)\\.")
	for i := int64(0); ; i++ {
		p.IntSet(i)
		resp, err := tp.Execute()
		if err != nil {
			panic(err)
		}
		etor := extractor.ExtractHtml(resp.Content())
		xps, err := etor.XPaths("//div[@class='iC-column']")
		if err != nil {
			panic(err)
		}
		lis, errs := xps.ForEach("//li")
		if len(errs) != 0 {
			panic(errs)
		}
		if lis == nil {
			log.Println("page:", i)
			break

		}

		ididomCollection := lis.ForEachObjectByTag(Ididom{})
		for _, ididom := range ididomCollection {
			i := ididom.(*Ididom)
			ididoms = append(ididoms, i)
		}
	}
	CompressData("ididom.gob.zst", ididoms)
}
