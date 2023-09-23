namespace go sourceService

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

service sourceService{
    AddSourceResp addSource(1:AddSourceReq addReq),
    FindSourceByNameResp findSourceByName(1:FindSourceByNameReq findReq),
    GetPageSourceResp getSourcesByPage(1:GetPageSourceReq getReq),
}

