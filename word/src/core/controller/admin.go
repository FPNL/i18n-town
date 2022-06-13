package controller

import (
	"fmt"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/igrpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IAdminHandler interface {
	Login(*gin.Context)
	Auth(*gin.Context)
	Register(*gin.Context)
	Ping(*gin.Context)
}

type adminHandler struct {
	adminService service.IAdminService
}

var singleAdmin = adminHandler{}

func Admin(adminService service.IAdminService) IAdminHandler {
	singleAdmin.adminService = adminService
	return &singleAdmin
}

func (a *adminHandler) Ping(ctx *gin.Context) {
	ping, err := a.adminService.Ping(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, ping)
}

func (a *adminHandler) Login(ctx *gin.Context) {
	var person igrpc.SimplePerson

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

	ctx.Header("token", token)
	ctx.String(http.StatusOK, "ok")
}

func (a *adminHandler) Auth(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	nickname, err := a.adminService.Validate(ctx, token)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "token not found")
		return
	}

	ctx.Status(http.StatusOK)
	fmt.Println("it's " + nickname)
}

func (a *adminHandler) Register(ctx *gin.Context) {
	person := igrpc.Person{}

	err := ctx.BindJSON(&person)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ok, err := a.adminService.Register(ctx, &person)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	} else if !ok {
		ctx.String(http.StatusInternalServerError, "不及格")
		return
	}

	ctx.String(http.StatusOK, "ok")
}
