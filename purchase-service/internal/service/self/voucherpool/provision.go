package voucherpool

import "context"

const (
	provisionTarget    = 200
	provisionMaxPerRun = 50
)

func (s *Service) Provision(ctx context.Context) error {
	repo := s.repo

	availableTotal, err := repo.GetAvailableTotal(ctx)
	if err != nil {
		return err
	}

	if availableTotal >= provisionTarget {
		return nil
	}

	quantity := provisionTarget - availableTotal
	if quantity > provisionMaxPerRun {
		quantity = provisionMaxPerRun
	}

	voucherCodes, err := s.GenVoucherForPool(ctx, quantity)
	if err != nil {
		return err
	}

	return repo.CreateBatchOrNothing(ctx, voucherCodes)
}
