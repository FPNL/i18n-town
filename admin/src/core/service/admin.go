package service

import (
	"context"
	"fmt"

	"github.com/FPNL/admin/src/core/entity"
	"github.com/FPNL/admin/src/core/repository"
	"github.com/FPNL/admin/src/lib/ierror"
	pb "github.com/FPNL/admin/src/lib/igrpc"
	"github.com/FPNL/admin/src/tool"

	"gorm.io/gorm"
)

var singleAdmin = Admin{}

func AdminService(adminRepo repository.IAdminRepository) pb.AdminServer {
	singleAdmin.adminRepo = adminRepo
	return &singleAdmin
}

type Admin struct {
	adminRepo repository.IAdminRepository
	pb.UnimplementedAdminServer
}

func (a *Admin) Ping(ctx context.Context, none *pb.None) (*pb.Pong, error) {
	return &pb.Pong{Ping: "PPPPPPPong"}, nil
}

func (a *Admin) Register(ctx context.Context, person *pb.RegisterInfo) (*pb.OK, error) {
	user := &entity.User{
		LoginInfo: entity.LoginInfo{
			Username: person.GetLoginInfo().GetUsername(),
			Password: person.GetLoginInfo().GetPassword(),
		},
		Role:     person.GetUser().GetRole(),
		Nickname: person.GetUser().GetNickname(),
		Organize: entity.Organize{
			Model: gorm.Model{
				ID: uint(person.GetUser().GetOrganize()),
			},
		},
	}

	if err := a.adminRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &pb.OK{Ok: true}, nil
}

func (a *Admin) Login(ctx context.Context, person *pb.LoginInfo) (*pb.Token, error) {
	loginInfo := &entity.LoginInfo{
		Username: person.GetUsername(),
		Password: person.GetPassword(),
	}

	// TODO what if user login twice?
	// 已經登入了為什麼還會登入一次，是否代表從不同代理/裝置/瀏覽器登入
	// code...

	user, err := a.adminRepo.FindUserBySimple(ctx, loginInfo)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, ierror.NewValidateErr("")
	}

	token, err := tool.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Pid: token}, nil
}

func (a *Admin) Authenticate(ctx context.Context, token *pb.Token) (*pb.User, error) {
	parseToken, err := tool.ParseToken(token.GetPid())
	if err != nil {
		return nil, err
	}

	id, ok := parseToken["_id"].(uint)
	if !ok {
		return nil, ierror.NewValidateErr(fmt.Sprintf("ID 來源有問題: %T %[1]v", parseToken["_id"]))
	}

	user, err := a.adminRepo.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.User{
			Organize: uint64(user.Organize.ID),
			Nickname: user.Nickname,
			Role:     user.Role,
		},
		nil
}
