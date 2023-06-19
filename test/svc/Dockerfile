FROM golang:latest 
RUN go get -v github.com/go-redis/redis
RUN go get -v github.com/gocql/gocql
RUN go get -v k8s.io/client-go/... 
RUN go get -v github.com/ant0ine/go-json-rest/rest
RUN go get -v github.com/pkg/errors
RUN apt-get update
RUN apt-get -y install vim
RUN apt-get install -y git-core
WORKDIR /go/src
RUN git clone https://github.com/humstarman/inClusterServiceDiscovery.git 
WORKDIR /workspace
ADD scripts /workspace/scripts/
ADD src /workspace/src/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/rest /workspace/src/rest.go
