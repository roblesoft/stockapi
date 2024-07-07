FROM golang:1.22.4 as builder

# building docker from stages

ENV PROJECT_DIR /go/bin/worker
WORKDIR $PROJECT_DIR

COPY ./go.* ./
RUN go mod download

COPY ./ ./

RUN go mod tidy
RUN go mod vendor

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# go build takes the result file and the second argument the source folder
RUN  go build -v -o app ./cmd

# use scratch to minimal image to excute the container
FROM scratch AS bin-unix

WORKDIR /app

# copy the result file from the build stage and copy into the workdir in the bin-unix stage
COPY --from=builder /go/bin/worker/app /app/server

EXPOSE 3005

# execute the binary
ENTRYPOINT ["./server"]
