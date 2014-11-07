package helper

import (
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/jamesyong/o3erp/go/thriftlib"
)

type ThriftFunc func(client *thriftlib.BaseServiceClient) (map[string]string, error)

func GetLoginFunction(userLoginId string, loginPwd string) ThriftFunc {
	return func(client *thriftlib.BaseServiceClient) (map[string]string, error) {
		return client.UserLogin(userLoginId, loginPwd)
	}
}

func GetHasPermissionFunction(userLoginId string, permissions []string) ThriftFunc {
	return func(client *thriftlib.BaseServiceClient) (map[string]string, error) {
		return client.HasPermission(userLoginId, permissions)
	}
}

func GetHasEntityPermissionFunction(userLoginId string, entities []string, actions []string) ThriftFunc {
	return func(client *thriftlib.BaseServiceClient) (map[string]string, error) {
		return client.HasEntityPermission(userLoginId, entities, actions)
	}
}

func GetMessageMapFunction(userLoginId string, resource string, labels []string) ThriftFunc {
	return func(client *thriftlib.BaseServiceClient) (map[string]string, error) {
		return client.GetMessageMap(userLoginId, resource, labels)
	}
}

func RunThriftService(thriftFunc ThriftFunc) (map[string]string, error) {

	//addr := flag.String("addr", "localhost:9090", "Address to listen to")
	//secure := flag.Bool("secure", false, "Use tls secure transport")
	//flag.Parse()
	addr := "localhost:9090"
	secure := true

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTTransportFactory()

	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return nil, err
	}
	client := thriftlib.NewBaseServiceClientFactory(transport, protocolFactory)

	return thriftFunc(client)
}
