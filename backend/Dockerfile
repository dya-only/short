FROM golang:alpine as builder
WORKDIR /build

RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=1 go build -o ./short
RUN chmod +x ./short

# ---

FROM alpine
WORKDIR /app

COPY --from=builder /build/short /app/short

ENTRYPOINT ["/app/short"]
