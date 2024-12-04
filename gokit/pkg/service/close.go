package service

import (
	"fmt"
	"go.uber.org/zap"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func (i *Impl) closeInternalServers() {
	logger := i.Logger()

	if len(i.internalServers) > 0 {
		for _, is := range i.internalServers {
			name := is.Name()
			logger.Info(fmt.Sprintf("Graceful shutdown for service %s starts", name))
			if err := is.Close(); err != nil {
				logger.Warn(fmt.Sprintf("Error when graceful shutdown for service %s", name), zap.Error(err))
			} else {
				logger.Info(fmt.Sprintf("Service %s is shut down successfully", name))
			}
		}
	}
}

func (i *Impl) finalizeInternalServers() {
	logger := i.Logger()

	if len(i.internalServers) > 0 {
		for _, is := range i.internalServers {
			name := is.Name()
			logger.Info(fmt.Sprintf("Finalize service %s starts", name))
			if err := is.Finalize(); err != nil {
				logger.Warn(fmt.Sprintf("Error when finalizing for service %s", name), zap.Error(err))
			} else {
				logger.Info(fmt.Sprintf("Service %s is finalized successfully", name))
			}
		}
	}

	if gcm := i.GRPCClientManager(); gcm != nil {
		logger.Info("Finalize gRPC client manager")
		gcm.Close()
	}
}

func (i *Impl) closeGRPC() {
	logger := i.Logger()

	if xdsGRPCServer := i.XDSGRPCServer(); xdsGRPCServer != nil {
		logger.Info("Shut down xDS gRPC server")
		xdsGRPCServer.GracefulStop()
	}

	if gRPCServer := i.GRPCServer(); gRPCServer != nil {
		logger.Info("Shut down gRPC server")
		gRPCServer.GracefulStop()
	}

	if lis := i.GRPCNetListener; lis != nil {
		logger.Info("Close gRPC net listener")
		lis.Close()
	}
}

func (i *Impl) closeGRPCGateway() {
	logger := i.Logger()

	GRPCGatewayClientConn := i.GRPCGatewayClientConn
	if GRPCGatewayClientConn != nil {
		logger.Info("Close gRPC gateway client connection")
		if err := GRPCGatewayClientConn.Close(); err != nil {
			logger.Warn("Error when closing gRPC gateway client connection", zap.Error(err))
		}
	}
}

func (i *Impl) closeHealthServer() {
	logger := i.Logger()

	healthServer := i.HealthServer()
	if healthServer != nil {
		healthServer.SetServingStatus("", healthv1.HealthCheckResponse_NOT_SERVING)
	}

	if healthServer != nil {
		logger.Info("Shut down health gRPC server")
		healthServer.Shutdown()
	}
}

func (i *Impl) stopCronjobManager() {
	logger := i.Logger()

	if cm := i.CronjobManager(); cm != nil {
		logger.Info("Stop cronjob manager")
		if err := cm.Stop(); err != nil {
			logger.Warn("Error when closing cronjob manager", zap.Error(err))
		}
	}
}

func (i *Impl) closeDB() {
	logger := i.Logger()

	if rueidisClient := i.Rueidis(); rueidisClient != nil {
		logger.Info("Shut down Rueidis client")
		rueidisClient.Close()
	}

	if db := i.MySQL(); db != nil {
		if sqlDB, err := db.DB(); err == nil {
			logger.Info("Shut down MySQL")
			if err = sqlDB.Close(); err != nil {
				logger.Warn("Error when graceful shutdown for MySQL", zap.Error(err))
			}
		}
	}
}

func (i *Impl) closeJWTManager() {
	logger := i.Logger()

	if jm := i.JWTManager(); jm != nil {
		logger.Info("Shut down JWT client manager")
		jm.Stop()
	}
}

func (i *Impl) Close() {
	i.state = Closing
	logger := i.Logger()

	if jm := i.JWTManager(); jm != nil {
		logger.Info("Shut down JWT client manager")
		jm.Stop()
	}

	i.closeHealthServer()
	i.stopCronjobManager()
	i.closeInternalServers()
	i.closeGRPCGateway()
	i.closeGRPC()
	i.closeDB()
	i.closeJWTManager()
	i.finalizeInternalServers()

	ctxCancelFunc := i.ctxCancelFunc
	if ctxCancelFunc != nil {
		logger.Info("Cancel context")
		i.ctxCancelFunc()
	}

	if i.AppConfig().Tracing.Enabled {
		logger.Info("Shut down tracer")
		tracer.Stop()
	}

	i.state = Closed
	logger.Info("Skit is shut down successfully")
}
