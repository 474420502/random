package random

import (
	"encoding/binary"
	"log"
	"math/rand"
	"time"
)

type Random struct {
	*rand.Rand
	label string
	seed  interface{}
	aplha []byte
}

// New Create Random Object
func New(seedOrLabel ...interface{}) *Random {

	r := &Random{aplha: alpha}

	defer func() {
		if r.label == "" {
			log.Println("seed:", r.seed) // 默认打印seed
		} else {
			log.Println(r.label, "seed:", r.seed) // 默认打印seed
		}
	}()

	switch len(seedOrLabel) {
	case 0:
		r.seed = time.Now().UnixNano()
		r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		return r
	case 1:

		switch s := seedOrLabel[0].(type) {
		case string:
			r.label = s
			r.seed = time.Now().UnixNano()
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case int:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case uint:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case int64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case uint64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(r.seed.(uint64))))
		case rand.Source64:
			r.seed = s
			r.Rand = rand.New(s)
		case rand.Source:
			r.seed = s
			r.Rand = rand.New(s)
		default:
			log.Panicln(s)
		}

		return r
	case 2:
		var seed interface{}
		if s, ok := seedOrLabel[0].(string); ok {
			r.label = s
			seed = seedOrLabel[1]
		} else {
			r.label = seedOrLabel[1].(string)
			seed = seedOrLabel[0]
		}

		switch s := seed.(type) {
		case string:
			r.label = s
			r.seed = time.Now().UnixNano()
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case int:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case uint:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case int64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case uint64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(r.seed.(uint64))))
		case rand.Source64:
			r.seed = s
			r.Rand = rand.New(s)
		case rand.Source:
			r.seed = s
			r.Rand = rand.New(s)
		}

		return r

	default:
		panic("param error.  empty like New() or seed,label or label,seed or seed or label")
	}
}

// New Create Random Object , Not log Seed
func NewNoLog(seedOrLabel ...interface{}) *Random {
	r := &Random{aplha: alpha}

	switch len(seedOrLabel) {
	case 0:
		r.seed = time.Now().UnixNano()
		r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		return r
	case 1:

		switch s := seedOrLabel[0].(type) {
		case string:
			r.label = s
			r.seed = time.Now().UnixNano()
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case int:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case uint:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case int64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case uint64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(r.seed.(uint64))))
		case rand.Source64:
			r.seed = s
			r.Rand = rand.New(s)
		case rand.Source:
			r.seed = s
			r.Rand = rand.New(s)
		default:
			log.Panicln(s)
		}

		return r
	case 2:
		var seed interface{}
		if s, ok := seedOrLabel[0].(string); ok {
			r.label = s
			seed = seedOrLabel[1]
		} else {
			r.label = seedOrLabel[1].(string)
			seed = seedOrLabel[0]
		}

		switch s := seed.(type) {
		case string:
			r.label = s
			r.seed = time.Now().UnixNano()
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case int:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case uint:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(s)))
		case int64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(r.seed.(int64)))
		case uint64:
			r.seed = s
			r.Rand = rand.New(rand.NewSource(int64(r.seed.(uint64))))
		case rand.Source64:
			r.seed = s
			r.Rand = rand.New(s)
		case rand.Source:
			r.seed = s
			r.Rand = rand.New(s)
		}

		return r

	default:
		panic("param error.  empty like New() or seed,label or label,seed or seed or label")
	}
}

// func (rand *Random)
// Seed 返回种子seed的值
func (rand *Random) GetSeed() interface{} {
	return rand.seed
}

// Bool 返回种子seed的值
func (rand *Random) Bool() bool {
	return rand.Intn(2) == 1
}

// Probability percent概率返回true. eg percent = 0.34 -> 34%
func (rand *Random) Probability(percent float64) bool {
	return rand.Float64() < percent
}

// OneOf64n N份之一触发概率
func (rand *Random) OneOf64n(N uint64) bool {
	return rand.Uint64()%N == 0
}

// OneOf32N N份之一触发概率
func (rand *Random) OneOf32n(N uint32) bool {
	return rand.Uint32()%N == 0
}

// Bytes 随机字节数据 [min,max)
func (rand *Random) Bytes(min, max int64) []byte {
	size := rand.Int63n(max+1-min) + min
	var buf []byte = make([]byte, size+9)
	for i := int64(0); i < size; i += 8 {
		binary.PutUvarint(buf[i:], rand.Uint64())
	}
	return buf[0:size]
}

// Execute random execute func [min, max) times  [min,max)
func (rand *Random) Execute(min, max int, do func()) {
	N := rand.Intn(max+1-min) + min
	for i := 0; i < N; i++ {
		do()
	}
}

func (rand *Random) Extend() *Extend {
	return (*Extend)(rand)
}
