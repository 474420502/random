package random

import "sort"

type Country struct {
	ID        int    `json:"id"`       // 国家代号
	LocalName string `json:"local"`    // 本地名. 就中文 或者 ...根据数据而定
	Name      string `json:"name"`     // 英文名
	FullName  string `json:"fullname"` // 全称
	Alpha2    string `json:"alpha2"`   // 2字母代号
	Alpha3    string `json:"alpha3"`   // 3字母代号
	Code      int    `json:"code"`     // 数字代号
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

// Country 随机国家
func (ext *Extend) Country() *Country {
	return countryData[ext.Intn(len(countryData))]
}
