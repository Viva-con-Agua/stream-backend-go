# drops-backend
drops implementation with go and echo

## use
install docker-compose

## redis
run `docker-compose up -d redis`
## database
run `docker-compose up -d drops-database && mysql -u drops -pdrops -h 172.2.200.1 drops < drops-database.sql`
