package random

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/474420502/requests"
	"github.com/klauspost/compress/zstd"
)

var defaultDir = ".random_base_data"
var sourceUrlPrefix = "https://raw.githubusercontent.com/474420502/random_data/master/"

func CheckAndDecompress(githuburl string, obj interface{}) {

	if info, err := os.Stat(defaultDir); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(defaultDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	} else {
		if !info.IsDir() {
			panic(fmt.Errorf("%s file exists and is not dir", defaultDir))
		}
	}

	idx := strings.LastIndex(githuburl, "/")
	filename := defaultDir + "/" + githuburl[idx+1:]

	f, err := os.Open(filename)
	if err != nil {

		resp, err := requests.NewSession().Get(githuburl).Execute()
		if err != nil {
			panic(err)
		}

		log.Println(len(resp.Content()))

		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}

		if len(resp.Content()) <= 100 {
			panic("找不到该数据")
		}

		w := bufio.NewWriter(f)

		_, err = w.Write(resp.Content())
		if err != nil {
			panic(err)
		}

		err = w.Flush()
		if err != nil {
			panic(err)
		}

		if err := f.Close(); err != nil {
			panic(err)
		}

		f, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
	}

	dec, err := zstd.NewReader(f)
	if err != nil {
		panic(err)
	}

	err = gob.NewDecoder(dec).Decode(obj)
	if err != nil {
		panic(err)
	}
}

func CompressData(filepath string, data interface{}) {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}

	enc, err := zstd.NewWriter(f)
	if err != nil {
		panic(err)
	}

	genc := gob.NewEncoder(enc)
	if err != nil {
		panic(err)
	}

	err = genc.Encode(data)
	if err != nil {
		panic(err)
	}

	enc.Close()
	f.Close()
}
