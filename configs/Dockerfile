# Created by Petr Lozhkin

FROM alpine as runtime

# othervise: "x509: failed to load system roots and no roots provided"
RUN apk add --no-cache ca-certificates

FROM golang:alpine as build

# Set golang envirovement and update PATH
ENV GOPATH=/go
ENV PATH=${GOPATH}/bin:/usr/local/go/bin:${PATH}

WORKDIR /go/src/github.com/im7mortal/test_ci
COPY . .

RUN go install github.com/im7mortal/test_ci/cmd/server

FROM runtime

COPY --from=build /go/bin/server /bin/server

CMD ["server"]
