FROM golang

ENV GOTRACEBACK=all

WORKDIR /app/go/src/github.com/bogdan-user/go-ticketing-app 

COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY ./services/auth services/auth/
COPY pkg pkg

RUN go build -o auth services/auth/main.go

CMD [ "./auth" ]