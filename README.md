# drone-deis

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-deis/status.svg)](http://cloud.drone.io/drone-plugins/drone-deis)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/deis.svg)](https://microbadger.com/images/plugins/deis "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-deis?status.svg)](http://godoc.org/github.com/drone-plugins/drone-deis)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-deis)](https://goreportcard.com/report/github.com/drone-plugins/drone-deis)

> Warning: This plugin has not been migrated to Drone >= 0.5 yet, you are not able to use it with a current Drone version until somebody volunteers to update the plugin structure to the new format.

Drone plugin to deploy or update a project on Deis. For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-deis/).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-deis
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/deis .
```

## Usage

```console
docker run --rm \
  -e PLUGIN_CONTROLLER=deis.deisapps.com \
  -e PLUGIN_APP=awesome \
  -e PLUGIN_FORCE=true \
  -e PLUGIN_COMMIT=true \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/deis
```
