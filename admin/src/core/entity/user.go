package entity

import (
	pb "github.com/FPNL/admin/src/lib/igrpc"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	LoginInfo
	Organize Organize `gorm:"type:VARCHAR(20);uniqueIndex:unique_oun;"`
	Role     pb.Role  `gorm:"type:smallint"`
	Nickname string   `gorm:"type:VARCHAR(20);uniqueIndex:unique_oun;"`
}

type Organize struct {
	gorm.Model
	Name string
}

type LoginInfo struct {
	Username string `gorm:"type:VARCHAR(20);uniqueIndex:unique_oun;"`
	Password string `gorm:"type:VARCHAR(20)"`
}
