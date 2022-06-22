package repository

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/FPNL/admin/src/core/entity"
	"github.com/go-redis/redis/v9"
)

type IAdminRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	FindUserById(ctx context.Context, id uint) (*entity.User, error)
	FindUserBySimple(ctx context.Context, user *entity.LoginInfo) (*entity.User, error)
}

var singleAdmin = Admin{}

func AdminRepository(model *gorm.DB, cache redis.Cmdable) IAdminRepository {
	singleAdmin.cache = cache
	singleAdmin.model = model
	return &singleAdmin
}

type Admin struct {
	cache redis.Cmdable
	model *gorm.DB
}

func (a *Admin) CreateUser(ctx context.Context, user *entity.User) error {
	var foundUser entity.User
	tx := a.model.Where("username = ?", user.Username).First(&foundUser)

	if errors.As(tx.Error, &gorm.ErrRecordNotFound) {
		return a.model.Create(user).Error
	} else if foundUser.ID > 0 {
		return status.Errorf(codes.InvalidArgument, "使用者重複了")
	} else {
		return tx.Error
	}
}

func (a *Admin) FindUserBySimple(ctx context.Context, loginInfo *entity.LoginInfo) (*entity.User, error) {
	var user entity.User
	tx := a.model.Model(&entity.User{}).Where("username = ? AND password = ?", loginInfo.Username, loginInfo.Password).First(&user)
	return &user, tx.Error
}

func (a *Admin) FindUserById(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	tx := a.model.Model(&entity.User{}).Where("ID = ?", id).First(&user)
	return &user, tx.Error
}
