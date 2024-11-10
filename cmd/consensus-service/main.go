package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wanglilind/qqq/internal/consensus/service"
	"github.com/wanglilind/qqq/pkg/config"
	"github.com/wanglilind/qqq/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	// åå§åéç½?
	cfg, err := config.Load("consensus-service")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// åå§åæ¥å¿?
	logger := logger.NewLogger(cfg.LogLevel)
	defer logger.Sync()

	// åå»ºå±è¯æå¡å®ä¾
	consensusService := service.NewConsensusService(cfg)

	// å¯å¨P2Pç½ç»
	if err := consensusService.StartP2PNetwork(); err != nil {
		logger.Fatalf("Failed to start P2P network: %v", err)
	}

	// åå»ºgRPCæå¡å?
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.GrpcInterceptor()),
	)

	// æ³¨åæå¡
	service.RegisterConsensusServiceServer(grpcServer, consensusService)

	// å¯å¨gRPCæå¡å?
	listener, err := net.Listen("tcp", cfg.GrpcAddress)
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}

	// ä¼éå³é­
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logger.Info("Shutting down consensus service...")
		consensusService.StopP2PNetwork()
		grpcServer.GracefulStop()
		cancel()
	}()

	// å¯å¨æå¡
	logger.Infof("Starting consensus service on %s", cfg.GrpcAddress)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
} 
