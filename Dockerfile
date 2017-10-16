# YAS3BL
# To populate data from yas3bl.json and generate README.md dynamically:
#   docker run --rm -it -v $(pwd):/go/src/github.com/petermbenjamin/yas3bl yas3bl
FROM golang:latest
LABEL maintainer="Peter Benjamin <petermbenjamin@gmail.com>"
ENTRYPOINT [ "go run main.go" ]
