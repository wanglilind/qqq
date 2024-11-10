package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/wanglilind/qqq/internal/transaction/service"
	"github.com/wanglilind/qqq/pkg/config"
	"github.com/wanglilind/qqq/pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	// åå§åéç½?
	cfg, err := config.Load("transaction-service")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// åå§åæ¥å¿?
	logger := logger.NewLogger(cfg.LogLevel)
	defer logger.Sync()

	// åå»ºäº¤ææå¡å®ä¾
	transactionService := service.NewTransactionService(cfg)

	// åå»ºgRPCæå¡å?
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.GrpcInterceptor()),
	)

	// æ³¨åæå¡
	service.RegisterTransactionServiceServer(grpcServer, transactionService)

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
		logger.Info("Shutting down gRPC server...")
		grpcServer.GracefulStop()
		cancel()
	}()

	// å¯å¨æå¡
	logger.Infof("Starting transaction service on %s", cfg.GrpcAddress)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
} 
