all: dep install

docker-compose:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	/go/bin/dep ensure
	go install ./kfig
	/go/bin/kfig -admin http://kong:8001

dep:
	dep ensure

install:
	go install ./kfig
