FROM golang:1.18-alpine as builder

LABEL stage=build

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -ldflags="-s -w" -o /card-deck-api cmd/api/main.go

###

FROM scratch

COPY --from=builder /card-deck-api /card-deck-api

EXPOSE 8080

ENTRYPOINT ["/card-deck-api"]
