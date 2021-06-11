module learn-golang

go 1.16

require (
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/golang-module/carbon v1.3.7
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/prometheus/client_golang v1.10.0 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
