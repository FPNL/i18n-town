package model

import (
	"context"
	"fmt"
	"strconv"

	"github.com/FPNL/admin/src/core/entity"
	"github.com/FPNL/admin/src/lib/ierror"

	"github.com/jackc/pgx/v4/pgxpool"
)

type IAdminModel interface {
	Insert(ctx context.Context, user *entity.AMI) error
	SelectById(ctx context.Context, id int) (*entity.AMI, error)
	SelectByLogin(ctx context.Context, user *entity.AMI) (*entity.AMI, error)
	SelectByUsername(ctx context.Context, username string) (*entity.AMI, error)
}

var singleAdmin = admin{}

func AdminModel(db *pgxpool.Pool) IAdminModel {
	singleAdmin.db = db
	return &singleAdmin
}

type admin struct {
	db *pgxpool.Pool
}

const TABLE_AMI = "ami"

func (a *admin) Insert(ctx context.Context, user *entity.AMI) error {
	sentence := fmt.Sprintf(
		`INSERT INTO %s(organize, nickname, username, password) 
		VALUES ('%s', '%s', '%s', '%s');`,
		TABLE_AMI,
		user.Organize, user.Nickname, user.Username,
		user.Password,
	)

	exec, err := a.db.Exec(ctx, sentence)
	if err != nil {
		return err
	}

	if exec.RowsAffected() != 1 {
		return ierror.NewUnknownErr("寫入數量為" + strconv.FormatInt(exec.RowsAffected(), 10))
	}

	return nil
}

func (a *admin) SelectById(ctx context.Context, id int) (*entity.AMI, error) {
	sentence := fmt.Sprintf(`SELECT * FROM %s WHERE id = %d;`,
		TABLE_AMI, id)
	user := &entity.AMI{}
	err := a.db.QueryRow(ctx, sentence).Scan(&user.Id, &user.Organize, &user.Nickname,
		&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *admin) SelectByLogin(ctx context.Context, user *entity.AMI) (*entity.AMI, error) {
	sentence := fmt.Sprintf(`SELECT * FROM %s WHERE username = '%s' and password = '%s';`,
		TABLE_AMI,
		user.Username, user.Password)
	err := a.db.QueryRow(ctx, sentence).Scan(&user.Id, &user.Organize, &user.Nickname,
		&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *admin) SelectByUsername(ctx context.Context, username string) (*entity.AMI, error) {
	sentence := fmt.Sprintf(`SELECT * FROM %s WHERE username = '%s';`,
		TABLE_AMI, username)

	user := &entity.AMI{}
	err := a.db.QueryRow(ctx, sentence).Scan(&user.Id, &user.Organize, &user.Nickname,
		&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
