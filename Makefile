all: dep install

docker-compose:
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
	/go/bin/dep ensure
	go test ./kfig -v -args -admin http://kong:8001 -mock http://kong:8000

testing:
	go test ./kfig -args -admin http://localhost:8001 -mock http://localhost:8000

dep:
	dep ensure

install:
	go install ./kfig
