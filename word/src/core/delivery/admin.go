package delivery

import (
	"fmt"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/igrpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAdminDelivery interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Guard_authenticate(ctx *gin.Context)
}

type adminDelivery struct {
	adminService service.IAdminService
}

var singleAdmin = adminDelivery{}

func Admin(adminService service.IAdminService) IAdminDelivery {
	singleAdmin.adminService = adminService
	return &singleAdmin
}

func (a *adminDelivery) Login(ctx *gin.Context) {
	var person igrpc.LoginInfo

	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.String(http.StatusBadRequest, "who are you")
		return
	}

	token, err := a.adminService.Login(ctx, &person)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("pid", token, 0, "", "", false, true)
	ctx.String(http.StatusOK, "ok")
}

func (a *adminDelivery) Guard_authenticate(ctx *gin.Context) {
	token, err := ctx.Cookie("pid")
	if err != nil {
		ctx.String(http.StatusUnauthorized, "找不到 token: "+err.Error())
		ctx.Abort()
		return
	}

	user, err := a.adminService.Authenticate(ctx, token)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "grpc 找不到 token: "+err.Error())
		ctx.Abort()
		return
	}

	ctx.Set("user_from_admin_delivery", user)

	ctx.Next()
}

func (a *adminDelivery) Register(ctx *gin.Context) {
	person := igrpc.RegisterInfo{}

	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.String(http.StatusBadRequest, "註冊失敗，資料不完整: "+err.Error())
		return
	}
	fmt.Printf("%#v", &person)
	ok, err := a.adminService.Register(ctx, &person)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "有東西壞了", err.Error())
		return
	} else if !ok {
		ctx.String(http.StatusInternalServerError, "不及格")
		return
	}

	ctx.String(http.StatusOK, "ok")
}
