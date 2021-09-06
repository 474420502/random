package random

type Ididom struct {
	Name        string `exp:"./a/text()"`
	Description string `exp:"./div/p/text()"`
}

var ididoms []*Ididom

func init() {
	registers[DataIdidomChina] = func() {
		CheckAndDecompress(dconf.dataUrls["ididom"], &ididoms)
	}
}

func (ext *Extend) IDidom() *Ididom {
	return ididoms[ext.Intn(len(ididoms))]
}
