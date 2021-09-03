package random

import "sort"

var provinceData []string

func init() {
	registers[DataProvinceChina] = func() {

		CheckAndDecompress(dconf.dataUrls["province"], &provinceData)
		sort.Strings(provinceData)
	}
}

func (ext *Extend) Province() string {
	return provinceData[ext.Intn(len(provinceData))]
}
