# TODO: multi-stage builds
FROM golang:1.24.0-alpine AS builder

WORKDIR /build

# copy the source code into /build
COPY . .

# switch to build/backend after copy
WORKDIR ./backend

# build the executable into /boids
# using disabling CGO for static binaries due to alpine build
RUN CGO_ENABLED=0 go build -o boids

# open port 8080
EXPOSE 8080

# run the executable
ENTRYPOINT ["./boids"]

################################################################################