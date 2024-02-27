FROM golang:1.22-alpine AS firstbuild
WORKDIR /api_gateway
COPY ./ ./
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM scratch
COPY --from=firstbuild /api_gateway/build /api_gateway/
COPY --from=firstbuild /api_gateway/dev.env /api_gateway/
WORKDIR /api_gateway/
EXPOSE 3000
ENTRYPOINT [ "/api_gateway/build" ]
