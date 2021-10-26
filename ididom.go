package random

type Ididom struct {
	Name        string `exp:"./a/text()"`     // 成语
	Description string `exp:"./div/p/text()"` // 描述
}

var ididoms []*Ididom

func init() {
	registers[DataIdidomChina] = func() {
		CheckAndDecompress(dconf.dataUrls["ididom"], &ididoms)
	}
}

// Ididom 随机成语
func (ext *Extend) Ididom() *Ididom {
	return ididoms[ext.Intn(len(ididoms))]
}
