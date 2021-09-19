# Initial stage: download modules
FROM golang:1.17-alpine as builder

ENV config=docker

WORKDIR /app

COPY ./ /app

RUN go build ./cmd/rest-api

# ---
# Intermediate stage: Build the binary
FROM scratch as runner

COPY --from=builder ./app/rest-api ./app/rest-api

WORKDIR /app
ENV config=docker

EXPOSE 3000

CMD ["rest-api"]
