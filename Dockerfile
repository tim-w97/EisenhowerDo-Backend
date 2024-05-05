FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o main

# run the executable
CMD ["./main"]