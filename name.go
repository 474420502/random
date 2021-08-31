package random

var lastnamesurl = "https://raw.githubusercontent.com/474420502/random_data/master/lastnames.gob.zst"
var firstnamesurl = "https://raw.githubusercontent.com/474420502/random_data/master/firstnames.gob.zst"

var firstNameData []string = []string{}
var lastNameData []string = []string{}

func (ext *Extend) FirstName() string {
	return firstNameData[ext.rand.Intn(len(firstNameData))]
}

func (ext *Extend) LastName() string {
	return lastNameData[ext.rand.Intn(len(lastNameData))]
}

func (ext *Extend) FullName() string {
	return ext.FirstName() + ext.LastName()
}
