# plugin-kodo

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-s3-sync/status.svg)](http://beta.drone.io/drone-plugins/drone-s3-sync)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-s3-sync?status.svg)](http://godoc.org/github.com/drone-plugins/drone-s3-sync)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-s3-sync)](https://goreportcard.com/report/github.com/drone-plugins/drone-s3-sync)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)

Drone plugin to upload files with a KODO Bucket. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
make build
make image
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-kodo' not found or does not exist..
```

## Usage

Execute from the working directory:


```sh
docker run --rm \
  -e PLUGIN_HOST=https://upload.qbox.me
  -e PLUGIN_ACCESS_KEY=AK \
  -e PLUGIN_SECRET_KEY=SK \
  -e PLUGIN_BUCKET=release-candidates \
  -e PLUGIN_SOURCE=/bin/drone-kodo.md \
  -e PLUGIN_KEY=/data/drone-kodo \
  index.qiniu.com/spock/kodo-plugin
```
