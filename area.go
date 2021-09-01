package random

import "sort"

const (
	areaDataUrl_CN = "https://raw.githubusercontent.com/474420502/random_data/master/area.gob.zst"
)

var areaData []string

func init() {
	registers[DataAreaChina] = func() {
		CheckAndDecompress(areaDataUrl_CN, &areaData)
		sort.Strings(areaData)
	}
}

func (ext *Extend) Area() string {
	return areaData[ext.Intn(len(areaData))]
}
