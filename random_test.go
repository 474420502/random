package random

import (
	"log"
	"testing"
)

func TestIP(t *testing.T) {
	r := New()
	// Use(DataCityChina)
	for i := 0; i < 1000; i++ {
		log.Println(r.Extend().IPv6().String())
	}
}

func TestGetData(t *testing.T) {
	r := New()
	Use(DataCityChina)
	for i := 0; i < 1000; i++ {
		log.Println(r.Extend().City())
	}

}

func estOpenTxt(t *testing.T) {

	// var hostsuffix []string = []string{
	// 	"com",
	// 	"xyz",
	// 	"com",
	// 	"net",
	// 	"com",
	// 	"top",
	// 	"com",
	// 	"tech",
	// 	"com",
	// 	"org",
	// 	"com",
	// 	"gov",
	// 	"com",
	// 	"edu",
	// 	"com",
	// 	"ink",
	// 	"com",
	// 	"int",
	// 	"com",
	// 	"mil",
	// 	"com",
	// 	"pub",
	// 	"com",
	// 	"cn",
	// 	"com",
	// 	"com.cn",
	// 	"com",
	// 	"net.cn",
	// 	"com",
	// 	"gov.cn",
	// 	"com",
	// 	"org.cn",
	// }

	// rand := New()

	// f, _ := os.Open("word.txt")
	// data, _ := ioutil.ReadAll(f)
	// allwords := string(data)

	// var wordsmap map[string]bool = make(map[string]bool)
	// for _, s := range strings.Split(allwords, " ") {

	// 	s = strings.ToLower(s)

	// 	if len(s) < 3 {
	// 		continue
	// 	}

	// 	if len(s) > 8 {
	// 		continue
	// 	}

	// 	if rand.OneOf64n(3) {
	// 		num := rand.Intn(10000) + 100
	// 		ns := strconv.Itoa(num)
	// 		wordsmap[ns] = true
	// 	}

	// 	wordsmap[s] = true
	// }

	// var suffixwords [][]byte
	// for k := range wordsmap {
	// 	suffixwords = append(suffixwords, []byte("@"+k+"."+hostsuffix[rand.Intn(len(hostsuffix))]))
	// }

	// f, err := os.OpenFile("emailsuffix.gob.zst", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
	// if err != nil {
	// 	panic(err)
	// }

	// enc, err := zstd.NewWriter(f)
	// if err != nil {
	// 	panic(err)
	// }

	// genc := gob.NewEncoder(enc)
	// if err != nil {
	// 	panic(err)
	// }

	// err = genc.Encode(suffixwords)
	// if err != nil {
	// 	panic(err)
	// }

	// enc.Close()
	// f.Close()

}

var data []string

// func estWriteCase1(t *testing.T) {

// 	resp, err := gcurl.Parse("https://raw.githubusercontent.com/uiwjs/province-city-china/gh-pages/city.json").Temporary().Execute()
// 	if err != nil {
// 		panic(err)
// 	}
// 	for _, g := range gjson.Parse(string(resp.Content())).Array() {

// 		data = append(data, g.Get("name").String())
// 	}

// 	f, err := os.OpenFile("city.gob.zst", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
// 	if err != nil {
// 		panic(err)
// 	}

// 	enc, err := zstd.NewWriter(f)
// 	if err != nil {
// 		panic(err)
// 	}

// 	genc := gob.NewEncoder(enc)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = genc.Encode(data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	enc.Close()
// 	f.Close()
// }

// func estCase2(t *testing.T) {
// 	// zstd.NewWriter()
// 	f, err := os.Open("firstnames.gob")
// 	if err != nil {
// 		panic(err)
// 	}
// 	dec := gob.NewDecoder(f)
// 	var lastnames map[string]int
// 	err = dec.Decode(&lastnames)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// data, err := ioutil.ReadAll(f)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	f, err = os.OpenFile("lastnames.gob.zstd", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
// 	if err != nil {
// 		panic(err)
// 	}

// 	enc, err := zstd.NewWriter(f)
// 	if err != nil {
// 		panic(err)
// 	}

// 	gob.NewEncoder(enc).Encode(lastnames)

// 	// _, err = enc.Write(data)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// }
