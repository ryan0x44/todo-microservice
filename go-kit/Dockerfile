# build stage
FROM golang:1.9-alpine AS build
ADD . /go/src/github.com/ryan0x44/todo-microservice/go-kit
RUN go install github.com/ryan0x44/todo-microservice/go-kit/cmd/todo
# final stage
FROM golang:1.9-alpine
COPY --from=build /go/bin/todo /app/
ENTRYPOINT ["/app/todo"]

