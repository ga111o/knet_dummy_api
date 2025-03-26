FROM golang:1.21-alpine

WORKDIR /app

COPY go.* ./
COPY *.go ./

RUN go build -o main .

EXPOSE 60950

CMD ["./main"] 