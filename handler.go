package main

import (
	"context"
	baseresp "github.com/oneVegeDog/sourceNavigator/pkg/rpc/baseresp"
	userservice "github.com/oneVegeDog/sourceNavigator/pkg/rpc/userService"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userservice.CreateUserReq) (resp *baseresp.Baseresp, err error) {
	// TODO: Your code here...
	return
}

// FindUserByName implements the UserServiceImpl interface.
func (s *UserServiceImpl) FindUserByName(ctx context.Context, req *userservice.FindUserReq) (resp *userservice.FindUserResp, err error) {
	// TODO: Your code here...
	return
}

// CheckedUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckedUser(ctx context.Context, req *userservice.CheckUserReq) (resp *userservice.CheckUserResp, err error) {
	// TODO: Your code here...
	return
}
