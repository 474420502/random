package random

import (
	"sort"
)

var firstNameData []string
var lastNameData []string

func init() {
	registers[DataNameChina] = func() {

		var lastnames map[string]int
		CheckAndDecompress(dconf.dataUrls["firstnames"], &firstNameData)
		CheckAndDecompress(dconf.dataUrls["lastnames"], &lastnames)
		for k := range lastnames {
			lastNameData = append(lastNameData, k)
		}
		sort.Strings(firstNameData)
		sort.Strings(lastNameData)
	}
}

// FirstName ć§
func (ext *Extend) FirstName() string {
	return firstNameData[ext.Intn(len(firstNameData))]
}

// LastName ć
func (ext *Extend) LastName() string {
	ln := lastNameData[ext.Intn(len(lastNameData))]
	if ext.Int63n(10) == 0 {
		return string([]rune(ln)[:1])
	}
	return ln
}

// FullName ćšć
func (ext *Extend) FullName() string {
	return ext.FirstName() + ext.LastName()
}
