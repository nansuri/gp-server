FROM golang:1.16

WORKDIR /go/src/app
COPY . .

ENV GOPROXY=https://goproxy.io,direct
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /bin/myapp .

CMD ["/bin/myapp"] 
