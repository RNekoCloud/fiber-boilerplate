FROM golang:1.22.0-alpine

WORKDIR /simaku-elearning

COPY go.mod /simaku-elearning 
COPY . /simaku-elearning

RUN go mod tidy 
RUN go build -o /simaku-elearning/bin/main /simaku-elearning/main.go 

EXPOSE 8080

ENTRYPOINT ["/simaku-elearning/bin/main"]
CMD ["http-gw-srv", "--port", "8080"]

