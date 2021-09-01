package random

type RandomDataType int

const (
	_ RandomDataType = iota
	DataNameChina
	DataProvinceChina
	DataCityChina
	DataAreaChina
	DataTownChina
	DataCountryChina
	DataEmailChina
)

var registers map[RandomDataType]func() = make(map[RandomDataType]func())

func Use(t RandomDataType) {
	registers[t]()
}
