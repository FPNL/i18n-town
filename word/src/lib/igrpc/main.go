package igrpc

import (
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var conn *grpc.ClientConn
var client AdminClient

func Go() (conn *grpc.ClientConn, err error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err = grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("GRPC 不能連線: %v", err)
	}

	client = NewAdminClient(conn)

	return
}

func Connect() AdminClient {
	if client == nil {
		panic("架構錯誤")
	}
	return client
}
