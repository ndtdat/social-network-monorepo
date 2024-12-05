package singlepod

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

func DeleteSingPods(logger *zap.Logger, singlePodMap map[string]*Service) {
	for id, checker := range singlePodMap {
		if err := checker.MarkAsNotRunningIfPossible(context.TODO()); err != nil {
			logger.Error(fmt.Sprintf("Error when marking %s as not running", id))
		}
	}
}
