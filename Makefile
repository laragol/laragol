GOPATH := ${PWD}/vendor:${PWD}
export GOPATH

main: clean
	@echo "dev: Build for Development with on-file change recompilation"
	@echo "dev-mac: Build for Development on mac with on-file change recompilation"
	@echo "build-dev: Build for Development"
	@echo "build-dev-listen: Build for Development & Listen"
	@echo "deploy: Build static binary"

dev:
	@bin/reflex-linux --decoration="fancy" -s -r '.go' make build-dev-listen

dev-mac:
	@bin/reflex-mac --decoration="fancy" -s -r '.go' make build-dev-listen

build-dev:
	@go build -o bin/main src/main.go

build-dev-listen:
	@go build -o bin/main src/main.go && ./bin/main

deploy:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o bin/main src/main.go

install:
	@mkdir -p vendor
	@go get -u github.com/valyala/quicktemplate
	@go get -u github.com/valyala/fasthttp
	@go get -u github.com/qiangxue/fasthttp-routing
	@go get -u github.com/go-sql-driver/mysql
	@go get -u github.com/jinzhu/gorm

style-code:
	gofmt -s -w src

clean:
	@rm -f bin/main
