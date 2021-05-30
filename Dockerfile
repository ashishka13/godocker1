FROM golang:latest
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0 
RUN go mod vendor \
    && go build -o myapp .
EXPOSE 5005
ENTRYPOINT ["./myapp"]
