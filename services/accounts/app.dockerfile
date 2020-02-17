FROM golang:1.13.4-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR go/src/git.mip-consult.de/sde/suzuki-framework/microservices
COPY vendor ../../vendor
COPY accounts ./
RUN go build -o /go/bin/app ./cmd/accounts/main.go

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8888
CMD ["app"]
