package random

import (
	"fmt"
	"time"
)

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
	case string:
		st, err := time.Parse("2006-01-02", start.(string))
		if err != nil {
			panic(err)
		}
		et, err := time.Parse("2006-01-02", end.(string))
		if err != nil {
			panic(err)
		}
		s := st.UnixMicro()
		e := et.UnixMicro()
		return time.UnixMicro(s + ext.Int63n(e-s))
	default:
		panic(fmt.Errorf("数据类型错误 start:%T end:%T, 类型必须相同", start, end))
	}

}
