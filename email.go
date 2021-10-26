package random

import (
	"math/rand"
	"sort"
)

var alpha = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

var suffixwords [][]byte

func init() {
	registers[DataEmailChina] = func() {
		CheckAndDecompress(dconf.dataUrls["emailsuffix"], &suffixwords)
		sort.Slice(suffixwords, func(i, j int) bool {
			return len(suffixwords[i]) < len(suffixwords[j])
		})
	}
}

func (ext *Extend) getChars() []byte {
	size := rand.Intn(16) + 4
	var result []byte = make([]byte, 0, 32)
	for i := 0; i < size; i++ {
		rsel := ext.Intn(36)
		result = append(result, alpha[rsel])
	}
	return result
}

// Email 随机email
func (ext *Extend) Email() string {
	buf := ext.getChars()
	buf = append(buf, suffixwords[ext.Intn(len(suffixwords))]...)
	return string(buf)
}

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
