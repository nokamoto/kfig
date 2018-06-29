all: dep install

dep:
	dep ensure

install:
	go install ./kfig
