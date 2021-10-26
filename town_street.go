package random

import "sort"

var townData []string

func init() {
	registers[DataTownStreetChina] = func() {
		CheckAndDecompress(dconf.dataUrls["town"], &townData)
		sort.Strings(townData)
	}
}

// Town 随机乡镇街道 同 Street
func (ext *Extend) Town() string {
	return townData[ext.Intn(len(townData))]
}

// Street 随机乡镇街道 同 Town
func (ext *Extend) Street() string {
	return ext.Town()
}
