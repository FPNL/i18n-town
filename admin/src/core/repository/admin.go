package repository

import (
	"context"
	"github.com/FPNL/admin/src/core/entity"
	"github.com/FPNL/admin/src/lib/ierror"

	"github.com/FPNL/admin/src/core/model"
	"github.com/go-redis/redis/v9"
)

type IAdminRepository interface {
	CreateUser(ctx context.Context, user *entity.AMI) error
	FindUserById(ctx context.Context, id int) (*entity.AMI, error)
	FindUserBySimple(ctx context.Context, user *entity.AMI) (*entity.AMI, error)
}

var singleAdmin = Admin{}

func AdminRepository(adminModel model.IAdminModel, cache redis.Cmdable) IAdminRepository {
	singleAdmin.cache = cache
	singleAdmin.adminModel = adminModel
	return &singleAdmin
}

type Admin struct {
	cache      redis.Cmdable
	adminModel model.IAdminModel
}

func (a *Admin) CreateUser(ctx context.Context, user *entity.AMI) error {
	if user, err := a.adminModel.SelectByUsername(ctx, user.Username); err != nil {
		return err
	} else if user != nil {
		return ierror.NewValidateErr("user Duplicate")
	}
	return a.adminModel.Insert(ctx, user)
}

func (a *Admin) FindUserBySimple(ctx context.Context, user *entity.AMI) (*entity.AMI, error) {
	return a.adminModel.SelectByLogin(ctx, user)
}

func (a *Admin) FindUserById(ctx context.Context, id int) (*entity.AMI, error) {
	return a.adminModel.SelectById(ctx, id)
}
