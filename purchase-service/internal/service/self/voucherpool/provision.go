package voucherpool

import "context"

func (s *Service) Provision(ctx context.Context) error {
	var (
		repo               = s.repo
		provisionTarget    = s.serviceCfg.ProvisionVoucherCodeCron.Target
		provisionMaxPerRun = s.serviceCfg.ProvisionVoucherCodeCron.Target
	)

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
