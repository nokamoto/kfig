# Kfig

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
```

## Yaml
```yaml
# https://docs.konghq.com/0.13.x/admin-api/#consumer-object
consumers:
  - username: nokamoto
    custom_id: nokamoto
    present: yes
```