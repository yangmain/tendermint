package lite

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/types"
)

// ErrOldHeaderExpired means the old (trusted) header has expired according to
// the given trustingPeriod and current time. If so, the light client must be
// reset subjectively.
type ErrOldHeaderExpired struct {
	At  time.Time
	Now time.Time
}

func (e ErrOldHeaderExpired) Error() string {
	return fmt.Sprintf("old header has expired at %v (now: %v)", e.At, e.Now)
}

// ErrNewValSetCantBeTrusted means the new validator set cannot be trusted
// because < 1/3rd (+trustLevel+) of the old validator set has signed.
type ErrNewValSetCantBeTrusted struct {
	Reason types.ErrNotEnoughVotingPowerSigned
}

func (e ErrNewValSetCantBeTrusted) Error() string {
	return fmt.Sprintf("cant trust new val set: %v", e.Reason)
}

type ErrInvalidHeader struct {
	Reason string
}

func (e ErrInvalidHeader) Error() string {
	return fmt.Sprintf("invalid header: %v", e.Reason)
}
