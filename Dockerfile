FROM golang:alpine as build

# build stage
# create folder app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o avenger-clinic

# Final stage
FROM alpine
WORKDIR /app

COPY --from=build /app/avenger-clinic /app/avenger-clinic

ENTRYPOINT ["/app/avenger-clinic"]