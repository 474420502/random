package random

import "sort"

var cityData []string

func init() {
	registers[DataCityChina] = func() {
		CheckAndDecompress(dconf.dataUrls["city"], &cityData)
		sort.Strings(cityData)
	}
}

// City 随机城市
func (ext *Extend) City() string {
	return cityData[ext.Intn(len(cityData))]
}
