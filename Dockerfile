FROM golang:alpine as builder

RUN apk --no-cache add ca-certificates
RUN update-ca-certificates

WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
RUN go build -mod=vendor -o /src/bin/server cmd/server/main.go
RUN go build -mod=vendor -o /src/bin/migrate cmd/migrate/main.go

FROM scratch as run

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /src/bin/server /server
COPY --from=builder /src/templates /templates
ENTRYPOINT ["/server"]
EXPOSE 10000

FROM scratch as migrate
COPY --from=builder /src/bin/migrate /migrate
COPY --from=builder /src/internal/database/migrations /db/migrations
ENTRYPOINT [ "/migrate" ]
