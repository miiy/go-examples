package etcd

import (
	etcd_client "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {


}

func NewClient() {
	etcdClient, err := etcd_client.New(etcd_client.Config{
		Endpoints:            []string{"http://254.0.0.1:12345"},
		AutoSyncInterval:     0,
		DialTimeout:          2 * time.Second,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		Context:              nil,
		Logger:               nil,
		LogConfig:            nil,
		PermitWithoutStream:  false,
	})
	if err != nil {

	}
	defer etcdClient.Close()
}