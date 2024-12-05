package voucher

import "context"

func (s *Service) MonitorUserVoucher(ctx context.Context) error {
	return s.userVoucherRepo.MarkUnavailableForExpiredBatch(ctx)
}
