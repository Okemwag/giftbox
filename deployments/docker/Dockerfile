FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN go build -o /bin/giftbox-api ./cmd/api

FROM alpine:3.22
RUN adduser -D -H giftbox
USER giftbox
COPY --from=build /bin/giftbox-api /bin/giftbox-api
EXPOSE 8080
ENTRYPOINT ["/bin/giftbox-api"]
