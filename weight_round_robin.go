package balancer

import (
	"errors"
	"strconv"
)

type WeightRoundRobinBalance struct {
	curIndex int
	data     []*WeightNode
}

type WeightNode struct {
	addr            string
	weight          int
	currentWeight   int
	effectiveWeight int
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params len need 2")
	}
	weight, err := strconv.Atoi(params[1])
	if err != nil {
		return err
	}
	r.data = append(r.data, &WeightNode{addr: params[0], weight: weight, effectiveWeight: weight})
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	var best *WeightNode
	total := 0
	for i := 0; i < len(r.data); i++ {
		w := r.data[i]
		total += w.effectiveWeight
		w.currentWeight += w.effectiveWeight
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}

	if best == nil {
		return ""
	}
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get() (string, error) {
	return r.Next(), nil
}
