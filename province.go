package random

import "sort"

const (
	provinceDataUrl_CN = "https://raw.githubusercontent.com/474420502/random_data/master/province.gob.zst"
)

var provinceData []string

func init() {
	registers[DataProvinceChina] = func() {
		CheckAndDecompress(provinceDataUrl_CN, &provinceData)
		sort.Strings(provinceData)
	}
}

func (ext *Extend) Province() string {
	return provinceData[ext.Intn(len(provinceData))]
}
