# todo

- ore calc inside construction
- load expenses from wallet api (job cost, sell taxes)

# run dev
server

`go run ./server/app.go`

run & reset cache
`go run ./server/app.go --reset-cache`

run import

`go run ./sde/app.go --reset-cache --file sde.zip`

set variables in var.json by example

# db & migrations (MySQL)

`create database publica;`

`create user 'publica'@'localhost' identified by 'publica';`

`grant all privileges on publica.* to 'publica'@'localhost';`
