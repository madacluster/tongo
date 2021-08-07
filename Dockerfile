FROM golang:1.15-alpine AS build
WORKDIR /src
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags="-w -s" -o /out/langui main.go
FROM scratch
COPY --from=build /out/langui /langui
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV LANGUI_LOOP=10
ENV LANGUI_LOOP=5
ENTRYPOINT ["/langui"]
