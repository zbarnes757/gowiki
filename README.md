# gowiki
Simple Wiki with page generation. Hooked up to local CockroachDB

In order to setup:

1. Install [cockroachDB](https://www.cockroachlabs.com/docs/v2.0/install-cockroachdb)
2. Get an instance of the db running in a separate terminal.
```bash
cockroach start \
--insecure \
--store=hello-1 \
--host=localhost
```
3. Create a user.
```bash
cockroach user set maxroach --insecure
```
4. Create an instance of the DB.
```bash
cockroach sql --insecure -e 'CREATE DATABASE wiki'
```
5. Give the user access to the database.
```bash
cockroach sql --insecure -e 'GRANT ALL ON DATABASE wiki TO maxroach'
```
6. Start the app!
```bash
go run main.go
```
