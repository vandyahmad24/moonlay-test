

# Moonlay Test Code By Vandy Ahmad

## Before Run You Must Install :
1. Setting .env with your local enviroment
2. Install make file if you use linux already bundled,
   In Windows you can install with choco command :
   ``` choco install make ```
3. If you use Postgresql DB you must change .env in DB_DRIVER=PSQL default DB is MYSQL
4. run command go mod tidy

## How to run in local
1. Running Program with command
   ``` make run-api ```
2. Import Postman collection for test api endpoint

## How to run unit test
1. go to folder usecase cd/app/usecase
2. go test -v -cover