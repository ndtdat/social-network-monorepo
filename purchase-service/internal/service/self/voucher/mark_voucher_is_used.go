package voucher

import "gorm.io/gorm"

func (s *Service) MarkUserVoucherIsUsedWithTx(tx *gorm.DB, userVoucherID, voucherConfigurationID uint64) error {
	if err := s.userVoucherRepo.MarkVoucherIsUsed(tx, userVoucherID); err != nil {
		return err
	}

	if err := s.configurationRepo.IncreaseRedeemedQty(tx, voucherConfigurationID); err != nil {
		return err
	}

	return nil
}
