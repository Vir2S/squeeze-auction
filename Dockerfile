FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY ./cmd/ /app/cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/squeeze-auction ./cmd/main.go

FROM alpine:latest as production
WORKDIR /app
COPY --from=builder /app/squeeze-auction /app/squeeze-auction

EXPOSE 3000
CMD ["/app/squeeze-auction"]
