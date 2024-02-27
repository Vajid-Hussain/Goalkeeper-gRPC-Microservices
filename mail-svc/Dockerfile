FROM golang:1.22-alpine AS firstbuild
WORKDIR /mail_svc
COPY ./ ./
RUN go mod download
RUN go build -o ./build ./cmd/main.go

FROM scratch
COPY --from=firstbuild /mail_svc/build /mail_svc/
COPY --from=firstbuild /mail_svc/dev.env /mail_svc/
WORKDIR /mail_svc/
EXPOSE 3003
ENTRYPOINT [ "/mail_svc/build" ]
