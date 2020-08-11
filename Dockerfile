# BUILD
FROM golang:alpine as builder
LABEL maintainer="jamie.kolles@axelspringer.com"

# Install git + SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

COPY . /swerve
WORKDIR /swerve

### TODO re-add the go-get line when the golang/crypto PR is merged
### https://github.com/golang/crypto/pull/143/commits/9860d605a5042280652d0aeeddc1f24782619181
# RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-w -s -X github.com/axelspringer/swerve/api.githubHash=`git rev-parse --short HEAD`" -o swerve main.go

# RUNTIME
FROM debian:latest
LABEL maintainer="jamie.kolles@axelspringer.com"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /swerve/swerve /swerve

ENTRYPOINT [ "/swerve" ]
