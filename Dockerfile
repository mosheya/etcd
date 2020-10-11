FROM golang
ADD . /go/src/github.com/mosheya/etcd
ADD cmd/vendor /go/src/github.com/mosheya/etcd/vendor
RUN go install github.com/mosheya/etcd
EXPOSE 2379 2380
ENTRYPOINT ["etcd"]
