namespace java o3erp.thrift
namespace go thriftlib

service BaseService {
    map<string,string> userLogin (1:string loginName, 2:string loginPwd)
}
