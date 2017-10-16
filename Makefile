.PHONY: readme html

readme:
	go run main.go

html: readme
	open docs/index.html

docker:
	docker run --rm -it -v $(shell pwd):/go/src/github.com/petermbenjamin/yas3bl -w /go/src/github.com/petermbenjamin/yas3bl golang:latest make
