package voucher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ndtdat/social-network-monorepo/purchase-service/internal/model"
	"os"
)

func (s *Service) Init(ctx context.Context) error {
	var (
		voucherCfgs     []*model.VoucherConfiguration
		voucherCfgsPath = s.serviceCfg.InitFilePath.VoucherCfg
	)

	voucherCfgsFile, err := os.ReadFile(voucherCfgsPath)
	if err != nil {
		return fmt.Errorf("cannot load voucher cfgs due to: %v", err)
	}

	if err = json.Unmarshal(voucherCfgsFile, &voucherCfgs); err != nil {
		return fmt.Errorf("cannot parse voucher configuration of path %s due to: %v", voucherCfgsPath, err)
	}

	return s.configurationRepo.CreateBatchOrNothing(ctx, voucherCfgs)
}
