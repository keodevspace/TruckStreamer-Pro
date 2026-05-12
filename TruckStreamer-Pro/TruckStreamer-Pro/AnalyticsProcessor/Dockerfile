FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod init AnalyticsProcessor && go build -o main .
CMD ["./main"]