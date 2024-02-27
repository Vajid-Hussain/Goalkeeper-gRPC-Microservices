FROM golang:1.22-alpine AS firstbuild
WORKDIR /mail-svc
COPY ./ ./
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM scratch
COPY --from=firstbuild /mail-svc/build /service
WORKDIR /service
ENTRYPOINT [ "build" ]
