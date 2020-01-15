FROM golang

RUN go get github.com/labstack/echo/...
RUN go get github.com/gocraft/dbr/...
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/configor/... 

WORKDIR /app

ADD . /app

CMD ["go", "run", "server.go"]
