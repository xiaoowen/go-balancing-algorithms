package balancer

type Balance interface {
	Add(...string) error
	Get() (string, error)
}

type LBType int

const Random LBType = 1
const RoundRobin LBType = 2
const WeightRoundRobin LBType = 3
const Shuffle LBType = 4

var LBTypeName = map[LBType]string{
	Random:           "random",
	RoundRobin:       "round_robin",
	WeightRoundRobin: "weight_round_robin",
	Shuffle:          "shuffle",
}

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
