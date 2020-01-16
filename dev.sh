

resetDatabase() {
  docker-compose rm -s drops-database && sudo rm -R volumes/drops-database/mysql && docker-compose up -d drops-database && mysql -u drops -pdrops -h 172.2.200.1 drops < drops-database.sql;
}

resetDatabase
