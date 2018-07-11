# Kfig

[![CircleCI](https://circleci.com/gh/nokamoto/kfig/tree/master.svg?style=svg)](https://circleci.com/gh/nokamoto/kfig/tree/master)

Kfig is a cli tool to configure [kong](https://github.com/Kong/kong) with a yaml file.

## Usage
```
$ kfig -h
Usage of kfig:
  -admin string
    	a kong admin api (default "http://localhost:8001")
  -yaml string
    	a yaml configration file (default "default.yaml")
```

## Build
```
$ make
```

## Quickstart
```
$ docker-compose up -d
$ kfig
$ curl localhost:8001/consumers
$ curl localhost:8001/key-auths
$ curl localhost:8001/services
```

## Yaml
```yaml
# https://docs.konghq.com/0.13.x/admin-api/#consumer-object
consumers:
  - username: nokamoto
    custom_id: nokamoto
    present: yes

    # https://docs.konghq.com/plugins/key-authentication
    key_auths:
      - key: my-api-key
        present: yes

# https://docs.konghq.com/0.13.x/admin-api/#service-object
services:
  - name: mock
    url: http://mockbin.org
    present: yes

    # https://docs.konghq.com/0.13.x/admin-api/#route-object
    # `present` field unsupported (kong does not provide any user defined identifiers for the route entities.)
    routes:
      - hosts:
          - example.com
        protocols:
          - http

    # https://docs.konghq.com/0.13.x/admin-api/#plugin-object
    # `present` field unsupported (use `enabled` instead)
    plugins:
      - name: key-auth
        config:
          hide_credentials: yes
        enabled: yes
```