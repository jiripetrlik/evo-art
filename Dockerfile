FROM golang:1.15-alpine as builder

WORKDIR /evo-art
COPY . .

RUN go build

FROM alpine:latest

COPY --from=builder /evo-art/evo-art /evo-art
EXPOSE 8080

CMD /evo-art
