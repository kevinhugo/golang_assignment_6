package main

import (
	"assignment/service"
)

func main() {
	userSvc := service.NewUserService()
	userSvc.Register(&service.User{Nama: "budi"})
	userSvc.Register(&service.User{Nama: "ke"})
	userSvc.Register(&service.User{Nama: "vin"})
	userSvc.Register(&service.User{Nama: "hu"})
	userSvc.Register(&service.User{Nama: "go"})
	userSvc.GetUser()
}