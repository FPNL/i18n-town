package main

import (
	"github.com/FPNL/admin/src/core/model"
	"github.com/FPNL/admin/src/core/repository"
	"github.com/FPNL/admin/src/core/service"
	"github.com/FPNL/admin/src/lib/icache"
	"github.com/FPNL/admin/src/lib/idatabase"
	"github.com/FPNL/admin/src/lib/igrpc"
	"google.golang.org/grpc"
)

func main() {
	var err error
	err = icache.Go()
	if err != nil {
		panic(err)
	}
	defer icache.Close()

	err = idatabase.Go()
	if err != nil {
		panic(err)
	}
	defer idatabase.Close()

	db := idatabase.Connect()
	cache := icache.Connect()
	mod := model.AdminModel(db)
	repo := repository.AdminRepository(mod, cache)
	serv := service.AdminService(repo)
	s := grpc.NewServer()
	igrpc.RegisterAdminServer(s, serv)

	err = igrpc.Go(s)
	if err != nil {
		panic(err)
	}
}
