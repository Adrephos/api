FROM golang:1.21.5-alpine AS build

WORKDIR /api
COPY . /api

RUN go mod tidy
RUN go build -o /bin/api ./main.go

FROM alpine:latest

COPY --from=build /bin/api /bin/api

ENV PORT="8080"
EXPOSE 8080

CMD ["/bin/api"]
