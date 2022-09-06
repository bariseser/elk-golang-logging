FROM golang:1.18.2
WORKDIR /app
COPY . .
RUN go build -o app main.go
EXPOSE 9090
CMD [ "./app" ]