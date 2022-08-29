FROM golang:1.18-buster AS build

WORKDIR go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /user-service


## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /user-service /user-service

EXPOSE 8085

USER nonroot:nonroot

ENTRYPOINT ["/user-service"]
