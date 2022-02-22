package random

import (
	"math/rand"
)

const texts = int(rune('㐀'))
const texte = int(rune('鿯'))
const textmax = texte + 1 - texts

// TextN 文字随机. N个文字
func (ext *Extend) TextN(n int) string {
	var result []rune
	for i := 0; i < n; i++ {
		v := rand.Intn(textmax)
		result = append(result, rune(texts+v))
	}
	return string(result)
}

// TextN 文字随机. N个文字 [sNum, eNum)
func (ext *Extend) TextRangeN(min, max int) string {
	var result []rune
	n := min + ext.Intn(max+1-min)
	for i := 0; i < n; i++ {
		v := rand.Intn(textmax)
		result = append(result, rune(texts+v))
	}
	return string(result)
}
