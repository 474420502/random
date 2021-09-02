package random

import (
	"log"
	"testing"
)

func TestLoadIndustry(t *testing.T) {
	// f, err := os.Open("./industry.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// reader := bufio.NewReader(f)
	// // var ind *industryRoot = &industryRoot{}

	// // tr := treelist.New()

	// var clslist []*industryClassNode

	// for {

	// 	cls := &industryClassNode{}

	// 	code, _, err := reader.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	cls.Code = string(code)

	// 	key, _, err := reader.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if key[0] >= '0' && key[0] <= '9' {
	// 		var childcode []byte = make([]byte, len(key))
	// 		copy(childcode, key)
	// 		if string(childcode) == "7920" {
	// 			log.Println("")
	// 		}
	// 		key, _, err = reader.ReadLine()
	// 		if err == io.EOF {
	// 			panic("")
	// 		}

	// 		cls.Name = string(key)
	// 		clslist = append(clslist, cls)

	// 		clslist = append(clslist, &industryClassNode{Code: string(childcode), Name: cls.Name})
	// 	} else {
	// 		cls.Name = string(key)
	// 		clslist = append(clslist, cls)
	// 	}

	// }

	// CompressData("industry.gob.zst", clslist)

	// var c1, c2, c3 *IndustryClass
	// for _, c := range clslist {
	// 	// log.Println(c.Code, c.Name)
	// 	if len(c.Code) == 1 {
	// 		ind.Category = append(ind.Category, c)
	// 		c1 = c
	// 		continue
	// 	}

	// 	if len(c.Code) == 2 {
	// 		c1.Children = append(c1.Children, c)
	// 		c2 = c
	// 		continue
	// 	}

	// 	if len(c.Code) == 3 {
	// 		c2.Children = append(c2.Children, c)
	// 		c3 = c
	// 		continue
	// 	}

	// 	if len(c.Code) == 4 {
	// 		c3.Children = append(c3.Children, c)
	// 		continue
	// 	}
	// }

	Use(DataIndustryChina)
	r := New(1)

	log.Println(industryTree.get(4, 3, 2))

	for i := 0; i < 100; i++ {
		log.Println(r.Extend().Industry())
	}

	// traverse(industryTree.Category)
}

func traverse(ic []*industryClassNode) {
	for _, c := range ic {
		log.Println(c.Code, c.Name)
		traverse(c.children)
	}
}
