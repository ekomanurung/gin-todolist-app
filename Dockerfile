FROM golang:1.18-alpine AS builder

# run add git to install dependencies
RUN apk update && apk add --no-cache git

# set working directory inside container
WORKDIR /app

# copy mod & sum file inside /app directory
COPY go.mod go.sum ./

# download the module dependencies
RUN go mod download

# copy all go source files into image directory
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
