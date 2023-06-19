FROM golang:latest as build
RUN go get -v github.com/go-redis/redis
RUN go get -v k8s.io/client-go/...
ADD src/svc.go /workspace/src/
ADD src/batch.go /workspace/src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /workspace/bin/batch /workspace/src/batch.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /workspace/bin/svc /workspace/src/svc.go
FROM bash:latest 
COPY --from=build /workspace/bin/batch /usr/local/bin/
COPY --from=build /workspace/bin/svc /usr/local/bin/
CMD ["/usr/local/bin/batch.sh"]
