

resetDatabase() {
  docker-compose rm -s stream-database && sudo rm -R volumes/stream-database/mysql && docker-compose up -d stream-database;
}

resetDatabase
