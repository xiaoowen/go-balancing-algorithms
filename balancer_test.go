package balancer

import (
	"log"
	"testing"
)

var addrs = [][]string{
	{"addr_1", "30"},
	{"addr_2", "50"},
	{"addr_3", "40"},
	{"addr_4", "0"},
}

func TestBalancer(t *testing.T) {
	strategies := []LBType{
		Random, RoundRobin, WeightRoundRobin, Shuffle,
	}
	for _, strategy := range strategies {
		log.Printf("balancer_strategy: %s\n", LBTypeName[strategy])
		balancer := BalanceFactory(strategy)
		doBalance(balancer)
	}
}

func doBalance(balancer Balance) {
	for _, v := range addrs {
		if err := balancer.Add(v...); err != nil {
			continue
		}
	}
	ret := make(map[string]int)
	for i := 0; i < 1000; i++ {
		addr, err := balancer.Get()
		if err != nil || addr == "" {
			continue
		}
		if _, ok := ret[addr]; !ok {
			ret[addr] = 0
		}
		ret[addr]++
	}
	for k, v := range ret {
		log.Printf("%s => %d\n", k, v)
	}
}
