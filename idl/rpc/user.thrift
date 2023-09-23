namespace go userService

include "baseresp.thrift"

struct User{
    1:i64 id
    2:string name
    3:string password
}

struct CreateUserReq{
    1:required string name
    2:required string password
}

struct FindUserReq{
    1:required string name
}

struct FindUserResp{
    1:User user
    2:bool isFind
    3:baseresp.baseresp baseresp
}

struct CheckUserReq{
    1:required string name
    2:required string password
}

struct CheckUserResp{
    1:bool isExsit
    2:baseresp.baseresp baseresp
}

service UserService{
    baseresp.baseresp register(1:CreateUserReq req),
    FindUserResp findUserByName(1:FindUserReq req),
    CheckUserResp checkedUser(CheckUserReq req),
}