#!/bin/bash

N=0

until [ $N -ge 3 ]
do
    docker-compose run golang curl --retry 60 --retry-delay 3 --retry-connrefused http://kong:8001/status && break
    N=$[$N+1]
done