package random

type Poetry struct {
	Title   string // 标题
	Dynasty string // 朝代
	Writer  string // 作者
	Content string // 诗词内容
}

var poetrys []*Poetry

func init() {
	registers[DataPoetryChina] = func() {
		CheckAndDecompress(dconf.dataUrls["poetry"], &poetrys)
	}
}

// Poetry 随机诗词
func (ext *Extend) Poetry() *Poetry {
	return poetrys[ext.Intn(len(poetrys))]
}
