package voucherpool

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/set"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/util"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
)

func (s *Service) getNotExistedCodesForPool(totalCodes, existedCodes []string) []*model.VoucherPool {
	existedCodeSet := set.New[string]()
	for _, ec := range existedCodes {
		existedCodeSet.Add(ec)
	}

	// Remove allocated gift codes
	var (
		notExistedCodes []*model.VoucherPool
		timeAt          = uint64(util.CurrentUnix())
	)
	for _, codes := range totalCodes {
		if !existedCodeSet.Contains(codes) {
			notExistedCodes = append(notExistedCodes, &model.VoucherPool{
				Code:      codes,
				CreatedAt: timeAt,
				UpdatedAt: timeAt,
			})
			timeAt++
		}
	}

	return notExistedCodes
}
