package igrpc

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn
var clientAdmin AdminClient

func Go() (err error) {
	// Set up a connection to the server.
	dns := fmt.Sprintf("%s:%s",
		os.Getenv("IADMIN_HOST"),
		os.Getenv("IADMIN_PORT"),
	)

	conn, err = grpc.Dial(dns, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("GRPC 不能連線: %v", err)
	}

	clientAdmin = NewAdminClient(conn)
	ping, err := clientAdmin.Ping(context.Background(), &None{})
	if err != nil {
		return err
	}

	fmt.Println(ping.GetPing() + " grpc admin 成功")

	return
}

func ConnectAdmin() AdminClient {
	if clientAdmin == nil {
		log.Fatalln("架構錯誤")
	}
	return clientAdmin
}

func Close() {
	if conn == nil {
		log.Fatalln("架構錯誤")
	}
	conn.Close()
}
