# Initial stage: download modules
FROM golang:1.17-alpine as builder

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go build ./cmd/api-rest


# Intermediate stage: Build the binary
FROM scratch as runner

COPY --from=builder ./app/api-rest ./app/api-rest

WORKDIR /app
ENV config=docker

EXPOSE 3000

CMD ["api-rest"]
