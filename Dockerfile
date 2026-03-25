# -------- BUILD STAGE --------
FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# build binary clearly
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# -------- RUN STAGE --------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]