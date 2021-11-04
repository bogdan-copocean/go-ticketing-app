FROM golang

ENV GOTRACEBACK=all

WORKDIR /app/go/src/github.com/bogdan-user/go-ticketing-app 

COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY ./services/tickets services/tickets/
COPY pkg pkg

RUN go build -o tickets services/tickets/main.go

CMD [ "./tickets" ]