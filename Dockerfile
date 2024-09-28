FROM golang:1.23.1
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o crud_app cmd/main.go
EXPOSE 8080
CMD [ "./crud_app" ]
