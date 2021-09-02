package random

type RandomDataType int

const (
	_ RandomDataType = iota
	// DataNameChina 姓名数据
	DataNameChina
	// DataProvinceChina 省份
	DataProvinceChina
	// DataCityChina 城市
	DataCityChina
	// DataTownChina 县城
	DataAreaChina
	// DataTownChina 街道(乡) 数据
	DataTownChina
	// DataCountryChina 全球国家 数据
	DataCountryChina
	// DataEmailChina email 数据
	DataEmailChina
	// DataIndustryChina 中国行业 数据
	DataIndustryChina
)

var registers map[RandomDataType]func() = make(map[RandomDataType]func())

func Use(t RandomDataType) {
	registers[t]()
}
