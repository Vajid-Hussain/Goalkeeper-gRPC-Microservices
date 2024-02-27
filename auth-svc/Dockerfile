FROM golang:1.22-alpine AS firstbuild
WORKDIR /auth_svc
COPY ./ ./
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM scratch
COPY --from=firstbuild /auth_svc/build /auth_svc/
COPY --from=firstbuild /auth_svc/dev.env /auth_svc/
WORKDIR /auth_svc/
EXPOSE 3001
ENTRYPOINT [ "/auth_svc/build" ]
