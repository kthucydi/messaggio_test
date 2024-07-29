FROM golang:1.21 AS builder

WORKDIR /app

COPY ./build/binary/messaggio_test /app/messaggio1
COPY ./.env /app/.env
# RUN cat /app/cmd/messaggio_test/main.go
# RUN go mod tidy
# RUN go build -o messagio1 /app/cmd/messaggio_test/main.go
EXPOSE 8050
# FROM golang:1.21
# COPY --from=builder /app/messagio1 /app/messagio1
RUN chmod +x /app/messaggio1
CMD ["/app/messaggio1", "-m=up"]


# FROM golang:1.21 AS builder
# # FROM golang:1.21
# WORKDIR /app

# COPY . .
# RUN cat /app/cmd/messaggio_test/main.go
# RUN go mod tidy
# RUN go build -o messagio1 /app/cmd/messaggio_test/main.go
# EXPOSE 8050
# FROM golang:1.21
# COPY --from=builder /app/messagio1 /app/messagio1
# RUN chmod +x /app/messagio1
# CMD ["/app/messagio1"]

# FROM golang:1.21.5 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# COPY . .

# RUN go build -o ./main ./cmd/messaggio_test/main 
# CMD ["/app/main"]