package main

import (
	loggerx "github.com/b3liv3r/logger"
	trainerv1 "github.com/b3liv3r/protos-for-gym/gen/go/trainer"
	"github.com/b3liv3r/trainers-for-gym/config"
	"github.com/b3liv3r/trainers-for-gym/modules/db"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/repository"
	"github.com/b3liv3r/trainers-for-gym/modules/trainers/service"
	server "github.com/b3liv3r/trainers-for-gym/modules/trainers/trpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	appConf := config.MustLoadConfig(".env")

	logger := loggerx.InitLogger(appConf.Name, appConf.Production)

	sqlDB, err := db.NewSqlDB(logger, appConf.Db)
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	repo := repository.NewTrainersRepository(sqlDB)
	service := service.NewTrainerService(repo, logger)
	s := InitRPC(service)
	lis, err := net.Listen("tcp", appConf.GrpcServerPort)
	if err != nil {
		logger.Error("failed to listen:", zap.Error(err))
	}
	logger.Info("grpc server listening at", zap.Stringer("address", lis.Addr()))
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", zap.Error(err))
	}
}

func InitRPC(tservice service.Trainer) *grpc.Server {
	s := grpc.NewServer()
	trainerv1.RegisterTrainerServer(s, server.NewTrainerRPCServer(tservice))

	return s
}
