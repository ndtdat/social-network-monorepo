package voucher

import "context"

func (s *Service) MonitorVoucherConfiguration(ctx context.Context) error {
	return s.configurationRepo.MarkUnavailableForExpiredBatch(ctx)
}
