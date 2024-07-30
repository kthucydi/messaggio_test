FROM golang:1.21

LABEL maintainer="Anton Bryntsev"

WORKDIR /app

COPY messaggio_test/ .
# if build by Dockerfile - uncomment next string and comment before string
# COPY . .

RUN go build -mod vendor -o /app/messaggio1 cmd/messaggio_test/main.go

EXPOSE 8050

RUN chmod +x /app/messaggio1

# arguments "-m=up" - tune on migration, delete it if you dont need in DB migration 
CMD ["/app/messaggio1", "-m=up"]
