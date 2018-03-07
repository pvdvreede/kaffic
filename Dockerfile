FROM golang:1.10 as dev
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/golang/mock/gomock
RUN go get -u github.com/golang/mock/mockgen
