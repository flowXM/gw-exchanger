package main

import (
	"flag"
	"fmt"
	proto_exchange "github.com/flowXM/proto-exchange/exchange"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gw-exchanger/internal/config"
	grpc_service "gw-exchanger/internal/grpc"
	"gw-exchanger/pkg/logger"
	"net"
)

func main() {
	var configFile string
	var port uint

	flag.StringVar(&configFile, "c", "", "Config file location")
	flag.UintVar(&port, "p", 5001, "Service port")
	flag.Parse()

	if configFile != "" {
		logger.Log.Debug("Loading env from file", "file", configFile)
		err := godotenv.Load(configFile)
		if err != nil {
			logger.Log.Error("Error loading .env file", "error", err)
			panic(err)
		}
	}

	config.Cfg = config.NewConfig()

	logger.Log.Info("Starting service", "port", port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Log.Error("Error start listener:", "error", err)
		panic(err)
	}

	logger.Log.Info("Service started")

	grpcServer := grpc.NewServer()
	proto_exchange.RegisterExchangeServiceServer(grpcServer, &grpc_service.ExchangeServiceServer{})

	if err := grpcServer.Serve(listener); err != nil {
		logger.Log.Error("Error start service:", "error", err)
	}
}
