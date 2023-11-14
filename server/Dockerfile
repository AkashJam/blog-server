FROM golang:alpine as builder
WORKDIR /go/src/app
# To get and install packagees from github
RUN apk add git

# To ensure that the git module can install by finding the go mod file
ENV GO111MODULE=auto
# Get CompileDaemon for live reload in dev
RUN go get github.com/githubnemo/CompileDaemon

# COPY go.mod .
# COPY go.sum .

COPY . .

RUN go mod download

# RUN go build -o ./run .

ENTRYPOINT CompileDaemon -build="go build -o /go/src/app" -command="/go/src/app"

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]
