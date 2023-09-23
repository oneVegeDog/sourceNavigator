namespace go userService

include "baseresp.thrift"


struct Source{
    1:string url
    2:i64 id
    3:string name
    4:string uploader_id
}

struct AddSourceReq{
    1:Source source
}

struct AddSourceResp{
    1:baseresp.baseresp baseresp
}

struct FindSourceByNameReq{
    1:string name
}

struct FindSourceByNameResp{
    1:Source source
    2:baseresp.baseresp baseresp
}

struct GetPageSourceReq{
    1:i64 pageNum
}

struct GetPageSourceResp{
    1:list<Source> sources
}


struct User{
    1:i64 id
    2:string name
    3:string password
    4:i8 right
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

service SourceNavigatorService{
    baseresp.baseresp register(1:CreateUserReq req),
    FindUserResp findUserByName(1:FindUserReq req),
    CheckUserResp checkedUser(CheckUserReq req),
    AddSourceResp addSource(1:AddSourceReq addReq),
    FindSourceByNameResp findSourceByName(1:FindSourceByNameReq findReq),
    GetPageSourceResp getSourcesByPage(1:GetPageSourceReq getReq),
}