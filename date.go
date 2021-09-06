package random

import (
	"log"
	"time"
)

func init() {

}

func (ext *Extend) Date() time.Time {
	return time.UnixMicro(ext.Int63n(time.Now().UnixMicro()))
}

func (ext *Extend) DateRange(start, end interface{}) time.Time {

	switch start.(type) {
	case int:
		s := start.(int)
		e := end.(int)
		return time.UnixMicro(int64(s + ext.Intn(e-s)))
	case int64:
		s := start.(int64)
		e := end.(int64)
		return time.UnixMicro(s + ext.Int63n(e-s))
	case time.Time:
		s := start.(time.Time).UnixMicro()
		e := end.(time.Time).UnixMicro()
		return time.UnixMicro(s + ext.Int63n(e-s))
	default:
		log.Panicln("数据类型错误", start, end)
	}

}
