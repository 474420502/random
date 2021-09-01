package random

import "sort"

const (
	cityDataUrl_CN = "https://raw.githubusercontent.com/474420502/random_data/master/city.gob.zst"
)

var cityData []string

func init() {
	registers[DataCityChina] = func() {
		CheckAndDecompress(cityDataUrl_CN, &cityData)
		sort.Strings(cityData)
	}
}

func (ext *Extend) City() string {
	return cityData[ext.Intn(len(cityData))]
}
