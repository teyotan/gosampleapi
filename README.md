migrate create -ext sql -dir migrations/ -seq MIGRATION_NAME
migrate -path migrations/ -database "mysql://USERNAME:PASSWORD@tcp(HOST:PORT)/DBNAME" -verbose up 