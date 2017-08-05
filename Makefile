build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo

image:
	docker build --rm=true -t index.qiniu.com/spock/kodo-plugin .

push:
	docker push index.qiniu.com/spock/kodo-plugin

all: build image push
