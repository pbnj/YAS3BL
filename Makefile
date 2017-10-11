.PHONY: readme

readme:
	go run main.go

docker:
	docker run --rm -it -v $(shell pwd):/go/src/github.com/petermbenjamin/yas3bl -w /go/src/github.com/petermbenjamin/yas3bl golang:latest make
