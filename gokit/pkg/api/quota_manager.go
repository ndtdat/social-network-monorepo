package api

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"go.uber.org/zap"
	"math/rand"
	"sync"
)

type QuotaManager struct {
	lastBannedMap    map[int]int64
	logger           *zap.Logger
	nKey             int
	maxRandomAttempt int
	banDurationInSec int

	//sync.Mutex
	sync.RWMutex
}

func NewQuotaManager(logger *zap.Logger, nKey int, maxRandomAttempt, banDurationInSec int) *QuotaManager {
	return &QuotaManager{
		nKey:             nKey,
		lastBannedMap:    make(map[int]int64),
		logger:           logger,
		maxRandomAttempt: maxRandomAttempt,
		banDurationInSec: banDurationInSec,
	}
}

func (q *QuotaManager) MarkKeyBanned(idx int) {
	if idx >= q.nKey {
		return
	}

	q.Lock()
	q.lastBannedMap[idx] = util.CurrentUnix()
	q.Unlock()
}

func (q *QuotaManager) markKeyUnBanned(idx int) {
	q.Lock()
	q.logger.Info(fmt.Sprintf("API Key idx %d is unbanned", idx))
	delete(q.lastBannedMap, idx)
	q.Unlock()
}

func (q *QuotaManager) GetUsableIdx() (int, error) {
	nKey := q.nKey
	if nKey == 0 {
		return 0, fmt.Errorf("there is no API keys to use")
	}

	nAttempt := 0

	for nAttempt < q.maxRandomAttempt {
		idx := rand.Intn(nKey) //nolint:gosec

		q.RLock()
		lastBanned, existed := q.lastBannedMap[idx]
		q.RUnlock()
		if !existed {
			return idx, nil
		}

		if util.CurrentUnix()-lastBanned >= int64(q.banDurationInSec) {
			q.markKeyUnBanned(idx)

			return idx, nil
		}

		nAttempt++
	}

	return 0, fmt.Errorf("reach max random attempt for finding API key")
}
