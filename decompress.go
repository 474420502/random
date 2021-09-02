package random

import (
	"encoding/gob"
	"os"
	"strings"

	"github.com/474420502/requests"
	"github.com/klauspost/compress/zstd"
)

func CheckAndDecompress(githuburl string, obj interface{}) {

	idx := strings.LastIndex(githuburl, "/")
	filename := githuburl[idx+1:]

	f, err := os.Open(filename)
	if err != nil {
		resp, err := requests.NewSession().Get(githuburl).Execute()
		if err != nil {
			panic(err)
		}
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			panic(err)
		}

		if len(resp.Content()) <= 100 {
			panic("找不到该数据")
		}

		_, err = f.Write(resp.Content())
		if err != nil {
			panic(err)
		}

		f.Close()

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
