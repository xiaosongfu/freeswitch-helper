# 应用名称
APP=freeswitch-helper

# 编译和运行的入口文件
ENTRY_FILE=./cmd/main.go

# 测试服务器地址
TEST_SERVER=root@192.168.160.11:/xiaoi/freeswitch-helper

# -------------------------------------------

all: clean test

run:
	go run $(ENTRY_FILE) --log-level=debug --config-file=config.dev.yaml

test:
	go test ./... -v

build:
	go build -o $(APP) $(ENTRY_FILE)

# build-linux
bl:
	GOOS=linux GOARCH=amd64 go build -o $(APP) $(ENTRY_FILE)

# build-windows
bw:
	GOOS=windows GOARCH=amd64 go build -o $(APP).exe $(ENTRY_FILE)

clean:
	rm -rf $(APP)

# -------------------------------------------

c-app:
	scp $(APP) $(TEST_SERVER)

# -------------------------------------------

gen-api-docs:
	swag init --generalInfo $(ENTRY_FILE) --output ./docs/api
	rm ./docs/api/docs.go

serve-api-docs:
	http-server ./docs/api --port=12053

docker-api-docs:
	docker build -t freeswitch-helper-api-docs:latest ./docs/api
