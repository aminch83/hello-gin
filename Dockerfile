# use golang image as base image
FROM golang:1.21.4-alpine3.18 as build

# set working directory
WORKDIR /app

# copy all files from current directory to working directory
COPY . .

# build the application
RUN go build -o ./bin/hello-gin


# use alpine image as base image
FROM alpine:3.18.4

# set working directory
WORKDIR /app
ENV GIN_MODE=release

# copy binary from build image to working directory
COPY --from=build /app/bin/ .

# expose port 8080
EXPOSE 8080

# run the application
CMD ["./hello-gin"]
