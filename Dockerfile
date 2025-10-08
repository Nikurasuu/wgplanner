FROM golang:1.25 AS build
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /server ./cmd

FROM scratch
COPY --from=build /server /server
EXPOSE 8080
CMD ["/server"]