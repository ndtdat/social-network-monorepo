package subscriptionplan

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"os"
)

func (s *Service) Init(ctx context.Context) error {
	var (
		subscriptionPlans    []*model.SubscriptionPlan
		subscriptionPlanPath = s.serviceCfg.InitFilePath.SubscriptionPlan
	)

	subscriptionPlanFile, err := os.ReadFile(subscriptionPlanPath)
	if err != nil {
		return fmt.Errorf("cannot load subscription plan due to: %v", err)
	}

	if err = json.Unmarshal(subscriptionPlanFile, &subscriptionPlans); err != nil {
		return fmt.Errorf("cannot parse subscription plan of path %s due to: %v", subscriptionPlanPath, err)
	}

	return s.subscriptionPlanRepo.CreateBatchOrNothing(ctx, subscriptionPlans)
}
