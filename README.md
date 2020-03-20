# stream-backend
stream service implementation with go and echo

## use
install docker-compose

## redis
run `docker-compose up -d redis`
## database
run `docker-compose up -d stream-database`

## development

### 1. Install Golang 
Like here: https://itrig.de/index.php?/archives/2377-Installation-einer-aktuellen-Go-Version-auf-Ubuntu.html

### Install dependecies
```
go get github.com/go-playground/validator
go get github.com/go-redis/redis
go get github.com/go-sql-driver/mysql
go get github.com/google/uuid
go get github.com/jinzhu/configor
go get github.com/labstack/echo
go get github.com/labstack/echo-contrib/session
go get github.com/rbcervilla/redisstore
go get golang.org/x/crypto/bcrypt
```

### 3. Checkout stream-backend-go
git clone https://github.com/Viva-con-Agua/stream-backend-go.git

### 4. Run server
Start server wiht `go run server.go`

### 5. update nginx
Update IP to you local IP in develop-pool branch at `routes/nginx-pool/pool.upstream` and restart nginx-pool docker with `docker restart pool-nginx`