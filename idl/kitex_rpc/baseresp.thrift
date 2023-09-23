

namespace go baseresp

include "code.thrift"

struct baseresp{
    1:string message
    2:code.Code statue_code
    3:i64 service_time
}