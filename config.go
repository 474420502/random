package random

type dataConfig struct {
	dataUrls map[string]string
}

var dconf *dataConfig = func() *dataConfig {
	var conf = &dataConfig{dataUrls: map[string]string{}}
	conf.setSourcePrefix(sourceUrlPrefix)
	return conf
}()

func (dconf *dataConfig) setSourcePrefix(prefix string) {
	if prefix[len(prefix)-1] != '/' {
		prefix += "/"
	}

	dconf.dataUrls["area"] = prefix + "area.gob.zst"
	dconf.dataUrls["city"] = prefix + "city.gob.zst"
	dconf.dataUrls["country"] = prefix + "country.gob.zst"
	dconf.dataUrls["emailsuffix"] = prefix + "emailsuffix.gob.zst"
	dconf.dataUrls["industry"] = prefix + "industry.gob.zst"
	dconf.dataUrls["firstnames"] = prefix + "firstnames.gob.zst"
	dconf.dataUrls["lastnames"] = prefix + "lastnames.gob.zst"
	dconf.dataUrls["province"] = prefix + "province.gob.zst"
	dconf.dataUrls["town"] = prefix + "town.gob.zst"
	dconf.dataUrls["ididom"] = prefix + "ididom.gob.zst"
	dconf.dataUrls["poetry"] = prefix + "poetry.gob.zst"
}

func ConfigSourcePrefix(prefix string) {
	dconf.setSourcePrefix(prefix)
}
