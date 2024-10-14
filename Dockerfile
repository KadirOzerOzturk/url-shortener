FROM golang:1.22

WORKDIR /url-shortener

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /url-shortener/cmd/server/main.go ./cmd/server

EXPOSE 3000

CMD [ "/url-shortener/cmd/server/main.go" ]
