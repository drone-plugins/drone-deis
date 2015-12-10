# Docker image for the Drone Deis plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-deis
#     make deps build
#     docker build --rm=true -t plugins/drone-deis .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    openssh \
    curl \
    perl && \
  rm -rf /var/cache/apk/*

ADD drone-deis /bin/
ENTRYPOINT ["/bin/drone-deis"]
