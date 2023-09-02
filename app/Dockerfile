FROM golang:1.21.0-bookworm as builder

WORKDIR /build

COPY go.mod ./
RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main

FROM gcr.io/distroless/static-debian11:nonroot

WORKDIR /app

COPY --from=builder /build/main ./

USER nonroot

EXPOSE 8080

CMD [ "/app/main" ]