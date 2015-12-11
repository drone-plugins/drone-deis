# drone-deis

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-deis/status.svg)](http://beta.drone.io/drone-plugins/drone-deis)
[![](https://badge.imagelayers.io/plugins/drone-deis:latest.svg)](https://imagelayers.io/?images=plugins/drone-deis:latest 'Get your own badge on imagelayers.io')


Drone plugin for deploying to Deis

## Usage

```sh
./drone-deis <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "controller": "deis.deisapps.com",
        "app": "awesome",
        "force": true,
        "commit": true
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-deis .
```

### Example

```sh
docker run -i plugins/drone-deis <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "controller": "deis.deisapps.com",
        "app": "awesome",
        "force": true,
        "commit": true
    }
}
EOF
```
