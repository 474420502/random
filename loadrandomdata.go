package random

import "sort"

type RandomDataType int

const (
	_ RandomDataType = iota
	DataNameChinese
)

func Use(t RandomDataType) {
	switch t {
	case DataNameChinese: // 加载文字数据
		var lastnames map[string]int
		CheckAndDecompress(firstnamesurl, &firstNameData)
		CheckAndDecompress(lastnamesurl, &lastnames)
		for k := range lastnames {
			lastNameData = append(lastNameData, k)
		}

		sort.Strings(firstNameData)
		sort.Strings(lastNameData)
	}
}
