psql -h 127.0.0.1 -U postgres -f ./sql/1-create-dbs.sql
psql -h 127.0.0.1 -U postgres -d shop -f ./sql/2-init-schema.sql
