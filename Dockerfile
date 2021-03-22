FROM arm64v8/golang:1.14

WORKDIR /go/src/app
COPY . .

RUN git clone https://github.com/nansuri/gp-server.git
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o /bin/myapp .

CMD ["/bin/myapp"] 