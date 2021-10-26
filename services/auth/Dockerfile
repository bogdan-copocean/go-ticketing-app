FROM golang

WORKDIR /app/go/src/github.com/bogdan-user/go-ticketing-app 

COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY ./services/auth services/auth/

RUN go build -o auth services/auth/main.go

CMD ["./auth"]