namespace go userService

include "baseresp.thrift"

struct User{
    1:i64 id
    2:string name
    3:string password
}

struct CreateUserReq{
    1:required string name (api.form="username")
    2:required string password (api.form="password")
}

struct FindUserReq{
    1:required string name (api.path="name")
}

struct FindUserResp{
    1:User user
    2:baseresp.baseresp baseresp
}

struct CheckUserReq{
    1:required string name (api.form="username")
    2:required string password (api.form="password")
}

struct CheckUserResp{
    1:bool isExsit
    2:baseresp.baseresp baseresp
}

service UserService{
    baseresp.baseresp register(1:CreateUserReq req)(
        api.post = "/user/register"
    ),
    FindUserResp findUserByName(1:FindUserReq req)(
        api.get = "user/find/:name"
    ),
    CheckUserResp checkedUser(1:CheckUserReq req)(
        api.post = "user/check"
    ),
}