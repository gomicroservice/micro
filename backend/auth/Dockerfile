FROM containerize/dep AS builder
WORKDIR /go/src/github.com/gotoolkit/app
# COPY Gopkg.toml Gopkg.lock ./
# RUN dep ensure -vendor-only
COPY . .
RUN go install .

FROM alpine
RUN apk add --no-cache ca-certificates
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf
COPY --from=builder /go/bin/app /usr/local/bin/app
WORKDIR /home/app
CMD ["app"]