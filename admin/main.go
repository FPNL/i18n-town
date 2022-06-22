package main

import (
	"github.com/FPNL/admin/src/core/repository"
	"github.com/FPNL/admin/src/core/service"
	"github.com/FPNL/admin/src/lib/icache"
	"github.com/FPNL/admin/src/lib/idatabase"
	"github.com/FPNL/admin/src/lib/igrpc"
	"github.com/FPNL/admin/src/setting"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var err error

	err = setting.Go()
	if err != nil {
		log.Fatalf("setting 專案架構錯誤: %v", err)
	}

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
	repo := repository.AdminRepository(db, cache)
	serv := service.AdminService(repo)
	s := grpc.NewServer()
	igrpc.RegisterAdminServer(s, serv)

	err = igrpc.Go(s)
	if err != nil {
		panic(err)
	}
}
