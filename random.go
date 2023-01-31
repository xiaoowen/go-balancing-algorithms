package balancer

import (
	"errors"
	"math/rand"
	"time"
)

type RandomBalance struct {
	curIndex int
	data     []string
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	addr := params[0]
	r.data = append(r.data, addr)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.data) == 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	r.curIndex = rand.Intn(len(r.data))
	return r.data[r.curIndex]
}

func (r *RandomBalance) Get() (string, error) {
	return r.Next(), nil
}
