package random

import "sort"

type Country struct {
	ID        int    `json:"id"`
	LocalName string `json:"local"`
	Name      string `json:"name"`
	FullName  string `json:"fullname"`
	Alpha2    string `json:"alpha2"`
	Alpha3    string `json:"alpha3"`
	Code      int    `json:"code"`
}

var countryData []*Country

func init() {
	registers[DataCountryChina] = func() {

		CheckAndDecompress(dconf.dataUrls["country"], &countryData)
		sort.Slice(countryData, func(i, j int) bool {
			return countryData[i].ID < countryData[j].ID
		})
	}
}

func (ext *Extend) Country() *Country {
	return countryData[ext.Intn(len(countryData))]
}
