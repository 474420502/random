package random

import "sort"

var areaData []string

func init() {
	registers[DataAreaChina] = func() {
		CheckAndDecompress(dconf.dataUrls["area"], &areaData)
		sort.Strings(areaData)
	}
}

// Area 随机区域
func (ext *Extend) Area() string {
	return areaData[ext.Intn(len(areaData))]
}
