package service

import "fmt"

type User struct {
	Nama string
}

var registeredUserList []User

type UserService struct {
}

type UserIface interface {
	Register(u *User)
	GetUser()
}

func NewUserService() UserIface {
	return &UserService{}
}

func (u *UserService) Register(user *User){
	registeredUserList = append(registeredUserList, *user)
	// return user.Nama + " berhasil didaftarkan"
}

func (u *UserService) GetUser() {
	for key, value := range registeredUserList{
		fmt.Println(key+1,value)
	}
}
