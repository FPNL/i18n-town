package service

import (
	"context"
	pb "github.com/FPNL/i18n-town/src/lib/igrpc"
)

type IAdminService interface {
	Validate(context.Context, string) (string, error)
	Login(context.Context, *pb.SimplePerson) (string, error)
	Register(context.Context, *pb.Person) (bool, error)
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

func (service *adminService) Validate(ctx context.Context, pid string) (string, error) {
	r, err := service.adminClient.Validate(ctx, &pb.Token{Pid: pid})
	if err != nil {
		return "", err
	}

	return r.GetNickname(), nil
}

func (service *adminService) Login(ctx context.Context, person *pb.SimplePerson) (string, error) {
	r, err := service.adminClient.Login(ctx, person)
	if err != nil {
		return "", err
	}

	return r.GetPid(), nil
}
func (service *adminService) Register(ctx context.Context, person *pb.Person) (bool, error) {
	r, err := service.adminClient.Register(ctx, person)
	if err != nil {
		return false, err
	}

	return r.GetOk(), nil
}
