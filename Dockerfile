FROM golang:1.10-alpine

RUN apk update && \
    apk add curl git

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Install pstore credential helper
RUN curl -sL -o /usr/bin/pstore https://github.com/glassechidna/pstore/releases/download/1.4.0/pstore_linux_amd64
RUN chmod 0755 /usr/bin/pstore
ENTRYPOINT ["pstore", "exec", "--verbose", "--", "/go/bin/app"]