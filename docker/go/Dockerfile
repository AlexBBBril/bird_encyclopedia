FROM golang:latest
RUN mkdir /app \
    && git clone https://github.com/gorilla/mux.git /app \
    && go get github.com/gorilla/mux

WORKDIR /app

RUN ls /app

# RUN go build main.go

# ENTRYPOINT [ "/app/main" ]
EXPOSE 8088