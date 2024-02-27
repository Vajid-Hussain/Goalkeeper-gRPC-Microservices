FROM golang:1.22-alpine AS firstbuild
WORKDIR /api-gateway
COPY ./ ./
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM scratch
COPY --from=firstbuild /api-gateway/build /service
WORKDIR /service
ENTRYPOINT [ "build" ]
