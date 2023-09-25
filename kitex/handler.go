package main

import (
	"context"
	source_rpc "github.com/oneVegeDog/sourceNavigator/pkg/source_rpc"
)

// SourceNavigatorServiceImpl implements the last service interface defined in the IDL.
type SourceNavigatorServiceImpl struct{}

// Register implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) Register(ctx context.Context, req *source_rpc.CreateUserReq) (resp *source_rpc.Baseresp, err error) {
	// TODO: Your code here...
	r := source_rpc.NewBaseresp()
	r.SetMessage("注册成功")
	return r, nil
}

// FindUserByName implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) FindUserByName(ctx context.Context, req *source_rpc.FindUserReq) (resp *source_rpc.FindUserResp, err error) {
	// TODO: Your code here...
	return
}

// CheckedUser implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) CheckedUser(ctx context.Context, req *source_rpc.CheckUserReq) (resp *source_rpc.CheckUserResp, err error) {
	// TODO: Your code here...
	return
}

// AddSource implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) AddSource(ctx context.Context, addReq *source_rpc.AddSourceReq) (resp *source_rpc.AddSourceResp, err error) {
	// TODO: Your code here...
	return
}

// FindSourceByName implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) FindSourceByName(ctx context.Context, findReq *source_rpc.FindSourceByNameReq) (resp *source_rpc.FindSourceByNameResp, err error) {
	// TODO: Your code here...
	return
}

// GetSourcesByPage implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) GetSourcesByPage(ctx context.Context, getReq *source_rpc.GetPageSourceReq) (resp *source_rpc.GetPageSourceResp, err error) {
	// TODO: Your code here...
	return
}
