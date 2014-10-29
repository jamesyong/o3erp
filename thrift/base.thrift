namespace java o3erp.thrift
namespace go thriftlib

service BaseService {
    map<string,string> userLogin (1:string userLoginId, 2:string loginPwd)
	map<string,string> hasPermission (1:string userLoginId, 2:list<string> permissions )
	map<string,string> hasEntityPermission (1:string userLoginId, 2:list<string> entities, 3:list<string> actions )
}
