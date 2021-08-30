package random

import (
	"log"
	"math/rand"
	"time"
)

type Random struct {
	*rand.Rand
	label string
	seed  interface{}
}

func New(seedOrLabel ...interface{}) *Random {

	r := &Random{}
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

func NewNoLog(seedOrLabel ...interface{}) *Random {
	r := &Random{}

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

// Probability 返回成功概率. eg percent = 0.34 -> 34%
func (rand *Random) Probability(percent float64) bool {
	return rand.Float64() < percent
}

// OneOf64n N份之一触发概率
func (rand *Random) OneOf64n(N uint64) bool {
	return rand.Uint64()%N == 0
}

// OneOf64N N份之一触发概率
func (rand *Random) OneOf32n(N uint32) bool {
	return rand.Uint32()%N == 0
}
