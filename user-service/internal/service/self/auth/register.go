package auth

import "context"

func (s *Service) Register(ctx context.Context) error {
	// TODO: Check email is exist

	// TODO: Get CODE param and check it's existed

	// TODO: If available => Call purchase-service to allocate voucher => Can ignore if failed

	// TODO: Use transaction to update data

	return nil
}
