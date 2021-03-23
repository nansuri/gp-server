FROM arm64v8/golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /bin/myapp .

CMD ["/bin/myapp"] 