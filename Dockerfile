FROM golang:1.25 AS build
WORKDIR /app
COPY . .
RUN go build -o /server ./cmd

FROM scratch
COPY --from=build /server /server
EXPOSE 8080
CMD ["/server"]