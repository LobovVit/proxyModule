package server

import (
	"go.uber.org/zap"
	"proxymodule/utils/config"
)

func GrpcStart(log *zap.Logger, Config config.Cfg) {

	log.Info("GrpcStart", zap.String("key", "val"))
}
