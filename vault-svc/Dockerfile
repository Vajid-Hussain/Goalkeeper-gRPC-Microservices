FROM golang:1.22-alpine AS firstbuild
WORKDIR /vault_svc
COPY ./ ./
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM scratch
COPY --from=firstbuild /vault_svc/build /vault_svc/
COPY --from=firstbuild /vault_svc/dev.env /vault_svc/
# run mkdir vault_svc
WORKDIR /vault_svc/
EXPOSE 3002
ENTRYPOINT [ "/vault_svc/build" ]
        