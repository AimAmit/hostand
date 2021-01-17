package remote

import (
	"flag"
	"fmt"
	"github.com/aimamit/hostand/proxy-server/proto"
	"google.golang.org/grpc"
	"log"
	"os"
)

type ServerCall struct {
	Main proto.ProxyClient
}

func ClientInit() proto.ProxyClient {
	MainHost := os.Getenv("MAIN_HOST")
	serverAddress := flag.String("address", fmt.Sprintf("%s:50051", MainHost), "tcp")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting remote server: %v\n", err)
	}
	mainClient := proto.NewProxyClient(conn)
	return mainClient
}

