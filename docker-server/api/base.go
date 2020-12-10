package api

import (
	"github.com/aimamit/hostand/docker-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)
type GrpcManger struct {}

func GrpcServerInit(){
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("Error while opening port %v\n", err)
	}
	s := grpc.NewServer()
	proto.RegisterDockerServiceServer(s, &GrpcManger{})
	log.Fatalln(s.Serve(lis))
}