package random

import "sort"

const (
	firstNameDataUrl_CN = "https://raw.githubusercontent.com/474420502/random_data/master/firstnames.gob.zst"
	lastNameDataUrl_CN  = "https://raw.githubusercontent.com/474420502/random_data/master/lastnames.gob.zst"
)

var firstNameData []string
var lastNameData []string

func init() {
	registers[DataNameChina] = func() {
		var lastnames map[string]int
		CheckAndDecompress(firstNameDataUrl_CN, &firstNameData)
		CheckAndDecompress(lastNameDataUrl_CN, &lastnames)
		for k := range lastnames {
			lastNameData = append(lastNameData, k)
		}
		sort.Strings(firstNameData)
		sort.Strings(lastNameData)
	}
}

func (ext *Extend) FirstName() string {
	return firstNameData[ext.Intn(len(firstNameData))]
}

func (ext *Extend) LastName() string {
	return lastNameData[ext.Intn(len(lastNameData))]
}

func (ext *Extend) FullName() string {
	return ext.FirstName() + ext.LastName()
}
