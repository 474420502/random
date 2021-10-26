package random

import "sort"

var provinceData []string

func init() {
	registers[DataProvinceChina] = func() {

		CheckAndDecompress(dconf.dataUrls["province"], &provinceData)
		sort.Strings(provinceData)
	}
}

// Province 随机省份
func (ext *Extend) Province() string {
	return provinceData[ext.Intn(len(provinceData))]
}
