package balancer

type Balance interface {
	Name() string
	Add(...string) error
	Get() (string, error)
}

type LBType int

const Random LBType = 1
const RoundRobin LBType = 2
const WeightRoundRobin LBType = 3
const Shuffle LBType = 4

func BalanceFactory(lbType LBType) Balance {
	switch lbType {
	case RoundRobin:
		return &RoundRobinBalance{}
	case WeightRoundRobin:
		return &WeightRoundRobinBalance{}
	case Shuffle:
		return &ShuffleBalance{}
	case Random:
	default:
		return &RandomBalance{}
	}
	return &RandomBalance{}
}
