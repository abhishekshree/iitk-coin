FROM golang:1.16.6

LABEL maintainer="Abhishek <abhishek.shree@outlook.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT ":3000"

RUN go build 

CMD ["./iitk-coin"]