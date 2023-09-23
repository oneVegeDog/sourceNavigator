package main

import (
	"context"
	sourcenavigatorservice "github.com/oneVegeDog/sourceNavigator/pkg/rpc/sourceNavigatorService/sourcenavigatorservice"
)

// SourceNavigatorServiceImpl implements the last service interface defined in the IDL.
type SourceNavigatorServiceImpl struct{}

// Register implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) Register(ctx context.Context, req *sourcenavigatorservice.CreateUserReq) (resp *sourcenavigatorservice.Baseresp, err error) {
	// TODO: Your code here...

	return
}

// FindUserByName implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) FindUserByName(ctx context.Context, req *sourcenavigatorservice.FindUserReq) (resp *sourcenavigatorservice.FindUserResp, err error) {
	// TODO: Your code here...
	return
}

// CheckedUser implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) CheckedUser(ctx context.Context, req *sourcenavigatorservice.CheckUserReq) (resp *sourcenavigatorservice.CheckUserResp, err error) {
	// TODO: Your code here...
	return
}

// AddSource implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) AddSource(ctx context.Context, addReq *sourcenavigatorservice.AddSourceReq) (resp *sourcenavigatorservice.AddSourceResp, err error) {
	// TODO: Your code here...
	return
}

// FindSourceByName implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) FindSourceByName(ctx context.Context, findReq *sourcenavigatorservice.FindSourceByNameReq) (resp *sourcenavigatorservice.FindSourceByNameResp, err error) {
	// TODO: Your code here...
	return
}

// GetSourcesByPage implements the SourceNavigatorServiceImpl interface.
func (s *SourceNavigatorServiceImpl) GetSourcesByPage(ctx context.Context, getReq *sourcenavigatorservice.GetPageSourceReq) (resp *sourcenavigatorservice.GetPageSourceResp, err error) {
	// TODO: Your code here...
	return
}
