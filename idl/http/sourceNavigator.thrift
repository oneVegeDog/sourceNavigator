namespace go http

enum Code{
    SUCCESS = 200;
    SERVER_BUSY = 500;
    NO_POWER = 401;
}

struct baseresp{
    1:string message
    2:Code statue_code
    3:i64 service_time
}

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
    2:baseresp baseresp
}

struct CheckUserReq{
    1:required string name (api.form="username")
    2:required string password (api.form="password")
}

struct CheckUserResp{
    1:bool isExsit
    2:baseresp baseresp
}

struct Source{
    1:string url
    2:i64 id
    3:string name
    4:string uploader_id
}

struct AddSourceReq{
    1:Source source (body="source")
}

struct AddSourceResp{
    1:baseresp baseresp
}

struct FindSourceByNameReq{
    1:string name (api.path="sourceName")
}

struct FindSourceByNameResp{
    1:Source source
    2:baseresp baseresp
}

struct GetPageSourceReq{
    1:i64 pageNum (api.path="pageNum")
}

struct GetPageSourceResp{
    1:list<Source> sources
}

service UserService{
    baseresp register(1:CreateUserReq req)(
        api.post = "/user/register"
    ),
    FindUserResp findUserByName(1:FindUserReq req)(
        api.get = "user/find/:name"
    ),
    CheckUserResp checkedUser(1:CheckUserReq req)(
        api.post = "user/check"
    ),
    AddSourceResp addSource(1:AddSourceReq addReq)(
        api.post = "/source/add"
    ),
    FindSourceByNameResp findSourceByName(1:FindSourceByNameReq findReq)(
        api.get = "/source/find/:sourceName"
    ),
    GetPageSourceResp getSourcesByPage(1:GetPageSourceReq getReq)(
        api.get = "/source/getPage/:pageNum"
    ),
}