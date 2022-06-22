package service

import (
	"context"
	pb "github.com/FPNL/i18n-town/src/lib/igrpc"
)

type IAdminService interface {
	Authenticate(context.Context, string) (*pb.User, error)
	Login(context.Context, *pb.LoginInfo) (string, error)
	Register(context.Context, *pb.RegisterInfo) (bool, error)
	Ping(ctx context.Context) (string, error)
}

type adminService struct {
	adminClient pb.AdminClient
}

var singleAdmin = adminService{}

func Admin(adminClient pb.AdminClient) IAdminService {
	singleAdmin.adminClient = adminClient
	return &singleAdmin
}

func (service *adminService) Ping(ctx context.Context) (string, error) {
	ping, err := service.adminClient.Ping(ctx, &pb.None{})
	if err != nil {
		return "", err
	}

	return ping.GetPing(), nil
}

func (service *adminService) Authenticate(ctx context.Context, pid string) (*pb.User, error) {
	r, err := service.adminClient.Authenticate(ctx, &pb.Token{Pid: pid})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (service *adminService) Login(ctx context.Context, person *pb.LoginInfo) (string, error) {
	r, err := service.adminClient.Login(ctx, person)
	if err != nil {
		return "", err
	}

	return r.GetPid(), nil
}

func (service *adminService) Register(ctx context.Context, person *pb.RegisterInfo) (bool, error) {
	r, err := service.adminClient.Register(ctx, person)
	if err != nil {
		return false, err
	}

	return r.GetOk(), nil
}
