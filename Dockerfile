# build
FROM golang:alpine as build
WORKDIR /go/src/ctrlshiftv
COPY . .
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
RUN go get -d -v ./...
RUN go install -v ./...

# run
FROM busybox:musl
WORKDIR /go/bin/
RUN chown -R 5000:5000 /go
RUN adduser -D -u 5000 app
USER app:app
COPY --from=build /go/bin/ctrlshiftv /go/bin/ctrlshiftv
EXPOSE 8000
CMD ["/go/bin/ctrlshiftv"]
