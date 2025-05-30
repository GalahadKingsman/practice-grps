package app

import (
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/app/userservice"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"github.com/GalahadKingsman/messenger_users/pkg/messenger_users_api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run(config *config.Configs) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	service := userservice.New()
	messenger_users_api.RegisterUserServiceServer(grpcServer, service)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println(err)
	}
}
