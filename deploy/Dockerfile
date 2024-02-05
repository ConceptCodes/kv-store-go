FROM golang:1.21.6-alpine AS base
WORKDIR /app

# builder
FROM base AS builder
ENV GOOS linux
ENV GOARCH amd64

# build-args
ARG VERSION

RUN apk --no-cache add bash git openssh

# modules: utilize build cache
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

# inject versioning information & build the binary
RUN go build -o app  .

# runner
FROM base AS runner
RUN apk add --no-cache libc6-compat tini
# Tini is now available at /sbin/tini

COPY --from=builder /app/app /app/app
EXPOSE 8080

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/app/app" ]