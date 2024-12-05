package campaign

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/user-service/internal/model"
	"os"
)

func (s *Service) Init(ctx context.Context) error {
	var (
		campaigns        []*model.Campaign
		campaignFilePath = s.serviceCfg.InitFilePath.Campaign
	)

	campaignFile, err := os.ReadFile(campaignFilePath)
	if err != nil {
		return fmt.Errorf("cannot load campaign due to: %v", err)
	}

	if err = json.Unmarshal(campaignFile, &campaigns); err != nil {
		return fmt.Errorf("cannot parse campaign of path %s due to: %v", campaignFilePath, err)
	}

	return s.repo.CreateBatchOrNothing(ctx, campaigns)
}
