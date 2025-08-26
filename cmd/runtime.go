package main

import (
	"exec-runtime/internal/backend"
	"exec-runtime/internal/imageapi"
	"exec-runtime/internal/runtimeapi"
	"exec-runtime/pkg/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	k8sCriApi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"net"
	"os"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	var network, address string
	if strings.HasPrefix(cfg.Listen, "unix://") {
		network = "unix"
		address = strings.TrimPrefix(cfg.Listen, "unix://")
		_ = os.Remove(address)
	} else if strings.HasPrefix(cfg.Listen, "tcp://") {
		network = "tcp"
		address = strings.TrimPrefix(cfg.Listen, "tcp://")
	} else {
		logger.Fatal("unsupported listen scheme", zap.String("listen", cfg.Listen))
	}

	lis, err := net.Listen(network, address)
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	grpcServer := grpc.NewServer()

	imageBackend, err := backend.BuildImageBackend(cfg.Image)
	if err != nil {
		panic(err)
	}
	imageService := imageapi.NewService(cfg, imageBackend, logger)
	runtimeService := runtimeapi.NewService(cfg, logger)

	k8sCriApi.RegisterRuntimeServiceServer(grpcServer, runtimeService)
	k8sCriApi.RegisterImageServiceServer(grpcServer, imageService)

	logger.Info("server started", zap.String("network", network), zap.String("address", address))
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("serve error", zap.Error(err))
	}
}
