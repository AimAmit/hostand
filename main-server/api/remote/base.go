package remote

import (
	"flag"
	"github.com/aimamit/hostand/main-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ProxyManger struct {}

type ClientCall struct {
	Docker proto.DockerServiceClient
}

var(
	InterServerCall = ClientCall{}
)

func Run() {
	serverAddress := flag.String("address", "0.0.0.0:50051", "tcp")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting remote server: %v\n", err)
	}
	dockerClient := proto.NewDockerServiceClient(conn)
	InterServerCall.Docker = dockerClient

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	s := grpc.NewServer()
	proto.RegisterProxyServer(s, &ProxyManger{})
	log.Println(s.Serve(lis))
}