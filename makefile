CTL = goctl
CORE_RPC_MODULES = user

# 生成api
gen-api:
	$(CTL) api go -api api/desc/all.api -dir ./api -style go_zero

# 生成rpc
gen-rpc:
	go run ./cmd/mergeProto.go -dir ./rpc/desc/core -out ./rpc/desc/core.proto
	$(CTL) rpc protoc ./rpc/desc/core.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc --style=go_zero

# 生成 ent
gen-ent:
	go generate ./rpc/ent

# 生成 swagger
gen-swagger:
	${CTL} api plugin -plugin goctl-swagger="swagger -filename core.json" -api ./api/desc/all.api -dir .

# 运行 api
run-api:
	go run ./api/core.go -f ./api/etc/core.yaml

# 运行 rpc
run-rpc:
	go run ./rpc/core.go -f ./rpc/etc/core.yaml