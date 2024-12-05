package campaign

import "context"

func (s *Service) Monitor(ctx context.Context) error {
	return s.repo.MarkUnavailableForExpiredCampaigns(ctx)
}
