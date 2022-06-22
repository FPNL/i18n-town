package igrpc

import (
	"fmt"
	"os"

	"log"
	"net"

	"google.golang.org/grpc"
)

func Go(s *grpc.Server) error {
	dns := fmt.Sprintf(":%s", os.Getenv("IADMIN_PORT"))
	lis, err := net.Listen("tcp", dns)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err.Error())
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err.Error())
	}

	return nil
}
