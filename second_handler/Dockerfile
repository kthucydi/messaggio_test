FROM golang:1.21

LABEL maintainer="Anton Bryntsev"

WORKDIR /app

COPY second_handler/ .
# if build by Dockerfile - uncomment next string and comment before string
# COPY . .
RUN go build -mod vendor -o /app/second_handler cmd/second_handler/main.go

EXPOSE 8050

RUN chmod +x /app/second_handler
CMD ["/app/second_handler"]
