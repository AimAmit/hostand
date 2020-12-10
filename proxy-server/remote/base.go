package remote

import (
	"flag"
	"github.com/aimamit/hostand/proxy-server/proto"
	"google.golang.org/grpc"
	"log"
)

type ServerCall struct {
	Main proto.ProxyClient
}

func ClientInit() proto.ProxyClient {
	serverAddress := flag.String("address", "0.0.0.0:50051", "tcp")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting remote server: %v\n", err)
	}
	mainClient := proto.NewProxyClient(conn)
	return mainClient
}

