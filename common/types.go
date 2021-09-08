package common

import (
	"math/big"
	"net/url"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/livepeer/go-livepeer/net"
)

type Broadcaster interface {
	Address() ethcommon.Address
	Sign([]byte) ([]byte, error)
}

type CapabilityComparator interface {
	CompatibleWith(*net.Capabilities) bool
	LegacyOnly() bool
}

const (
	Score_Untrusted = 0
	Score_Trusted   = 1
)

type OrchestratorLocalInfo struct {
	URL   *url.URL
	Score uint
}

type OrchestratorPool interface {
	GetInfos() []OrchestratorLocalInfo
	GetOrchestrators(int, Suspender, CapabilityComparator, uint) ([]*net.OrchestratorInfo, error)
	Size() int
	SizeWithScore(uint) int
}

type Suspender interface {
	Suspended(orch string) int
}

type OrchestratorStore interface {
	OrchCount(filter *DBOrchFilter) (int, error)
	SelectOrchs(filter *DBOrchFilter) ([]*DBOrch, error)
	UpdateOrch(orch *DBOrch) error
}

type RoundsManager interface {
	LastInitializedRound() *big.Int
}
