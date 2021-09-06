package random

import (
	"log"
	"strings"
	"testing"

	"github.com/474420502/extractor"
	"github.com/474420502/requests"
)

type cPoetry struct {
	Title   string   `exp:"./div[@class='pL-tit']//a[1]/text()"`
	Dynasty string   `exp:"./div[@class='pL-tit']//a[2]/text()"`
	Writer  string   `exp:"./div[@class='pL-tit']//a[3]/text()"`
	Content []string `exp:"./div[@class='cont']//div[@class='sj']/text()"`
}

func TestPoetry(t *testing.T) {
	r := New()
	Use(DataPoetryChina)
	for i := 0; i < 1000; i++ {
		p := r.Extend().Poetry()
		log.Println(p.Title)
		log.Println(p.Writer)
		log.Println(p.Dynasty)
		log.Println(p.Content)
	}
}

func estGetUrlPoetry(t *testing.T) {
	u := "https://www.y5000.com/poetry/list-0-0-1.html"
	tp := requests.NewSession().Get(u)
	p := tp.PathParam("list-(\\d+)-0-(\\d+)\\.")
	var ps []*cPoetry
	for c := int64(1); c < 13; c++ {
		for i := int64(1); ; i++ {
			p.IntArraySet(0, c)
			p.IntArraySet(1, i)

			resp, err := tp.Execute()
			if err != nil {
				panic(err)
			}
			etor := extractor.ExtractHtml(resp.Content())
			xps, err := etor.XPaths("//ul[@class='pL-ul']")
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

			for _, poetry := range lis.ForEachObjectByTag(cPoetry{}) {
				i := poetry.(*cPoetry)
				i.Dynasty = strings.Trim(i.Dynasty, "【】 ")
				ps = append(ps, i)
			}
		}
	}

	for _, cp := range ps {
		p := &Poetry{}
		p.Title = cp.Title
		p.Dynasty = cp.Dynasty
		p.Writer = cp.Writer
		p.Content = strings.Join(cp.Content, "\n")
		poetrys = append(poetrys, p)
	}

	CompressData("poetry.gob.zst", poetrys)
}
