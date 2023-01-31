package balancer

import (
	"errors"
	"math/rand"
	"time"
)

type ShuffleBalance struct {
	data []string
}

func (s *ShuffleBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	addr := params[0]
	s.data = append(s.data, addr)
	return nil
}

func (s *ShuffleBalance) Next() string {
	lens := len(s.data)
	if lens == 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	for i := lens - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	return s.data[0]
}

func (s *ShuffleBalance) Get() (string, error) {
	return s.Next(), nil
}
