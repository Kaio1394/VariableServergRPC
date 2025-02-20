package main

import (
	database "VariableServergRPC/internal/db"
	"VariableServergRPC/internal/handler"
	logger "VariableServergRPC/internal/logging"
	"VariableServergRPC/internal/repository"
	"VariableServergRPC/internal/service"
	"VariableServergRPC/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {

	logger.Log.Infoln("Connecting to database")
	db, err := database.ConnectDatabase()
	if err != nil {
		logger.Log.Errorf("Error connecting to database: %v", err)
		return
	}
	logger.Log.Infoln("Connect successfully")

	repo := repository.NewVariableRepositoryImpl(db)
	s := service.NewVariableServiceimpl(repo)
	h := handler.NewVariableHandlerImpl(s)

	logger.Log.Infoln("Starting server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Log.Errorf("Error listening on port 8080: %v", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterVariableServiceServer(server, h)
	if err := server.Serve(listener); err != nil {
		logger.Log.Errorf("Error serving on port 8080: %v", err)
		return
	}
	logger.Log.Infoln("Listening on port 8080")

}
