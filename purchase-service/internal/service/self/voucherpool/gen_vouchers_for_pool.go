package voucherpool

import (
	"context"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
)

const (
	VoucherCodeLength = 10
)

func (s *Service) GenVoucherForPool(ctx context.Context, qty int64) ([]*model.VoucherPool, error) {
	var (
		finalVoucherCodes []*model.VoucherPool
		remainQty         = qty
	)

	for {
		// Generate new gift codes
		generatingCodes, err := s.genVoucherCodes(VoucherCodeLength, remainQty)
		if err != nil {
			return nil, err
		}

		// Check existed in pool
		existedCodesOfPool, _, err := s.repo.GetExistedCodes(ctx, generatingCodes, false)
		if err != nil {
			return nil, err
		}

		generatingCodes = s.getNotExistedCodes(generatingCodes, existedCodesOfPool)

		// Check existed for new gift codes just created
		allocatedCodes, _, err := s.userVoucherRepo.GetAllocatedVoucherCodes(
			ctx, generatingCodes, false,
		)
		if err != nil {
			return nil, err
		}

		voucherCodesForPool := s.getNotExistedCodesForPool(generatingCodes, allocatedCodes)
		finalVoucherCodes = append(finalVoucherCodes, voucherCodesForPool...)
		nFinalVoucherCodes := int64(len(finalVoucherCodes))
		// Check gift codes are sufficient
		if nFinalVoucherCodes >= qty {
			break
		}

		// Continue if insufficient
		remainQty = qty - nFinalVoucherCodes
	}

	return finalVoucherCodes, nil
}
