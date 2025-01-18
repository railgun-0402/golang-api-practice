FROM golang:1.23-alpine

# /appを作業ディレクトリにする
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .

CMD ["go", "run", "main.go"]
