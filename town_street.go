package random

import "sort"

const (
	townDataUrl_CN = "https://raw.githubusercontent.com/474420502/random_data/master/town.gob.zst"
)

var townData []string

func init() {
	registers[DataTownChina] = func() {
		CheckAndDecompress(townDataUrl_CN, &townData)
		sort.Strings(townData)
	}
}

func (ext *Extend) Town() string {
	return townData[ext.Intn(len(townData))]
}

func (ext *Extend) Street() string {
	return ext.Town()
}
