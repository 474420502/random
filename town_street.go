package random

import "sort"

var townData []string

func init() {
	registers[DataTownStreetChina] = func() {
		CheckAndDecompress(dconf.dataUrls["town"], &townData)
		sort.Strings(townData)
	}
}

func (ext *Extend) Town() string {
	return townData[ext.Intn(len(townData))]
}

func (ext *Extend) Street() string {
	return ext.Town()
}
