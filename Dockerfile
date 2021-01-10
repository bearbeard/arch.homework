FROM golang:1.14 as builder
RUN mkdir -p /opt/application
WORKDIR /opt/application/
COPY . .
RUN go build -o app

FROM golang:1.14 as bundle
RUN mkdir -p /opt/application
WORKDIR /opt/application/
COPY --from=builder /opt/application/app ./app
EXPOSE 8000
ENTRYPOINT ["/bin/sh", "-c" , "exec ./app"]