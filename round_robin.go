package balancer

import "errors"

type RoundRobinBalance struct {
	curIndex int
	data     []string
}

func (r *RoundRobinBalance) Name() string {
	return "round_robin"
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	addr := params[0]
	r.data = append(r.data, addr)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	lens := len(r.data)
	if lens == 0 {
		return ""
	}
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	curAddr := r.data[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return curAddr
}

func (r *RoundRobinBalance) Get() (string, error) {
	return r.Next(), nil
}
