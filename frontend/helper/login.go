package helper

import (
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"o3erp/frontend/thriftlib"
)

func Login(loginName string, loginPwd string) (map[string]string, error) {

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

	return handleClient(client, loginName, loginPwd)
}

func handleClient(client *thriftlib.BaseServiceClient, loginName string, loginPwd string) (map[string]string, error) {
	return client.UserLogin(loginName, loginPwd)
}
